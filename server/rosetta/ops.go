package rosetta

import (
	"fmt"
	"github.com/coinbase/rosetta-sdk-go/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	abci "github.com/tendermint/tendermint/abci/types"

	util "github.com/cosmos/cosmos-sdk/server/rosetta/lib"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (c converter) ProcessMessage(msg sdk.Msg, events sdk.StringEvents, status string) ([]*types.Operation, error) {
	var txOps []*types.Operation
	var err error

	if util.Contains(StakingOperations, sdk.MsgTypeURL(msg)) || sdk.MsgTypeURL(msg) == MsgSendOperation {
		txOps, err = c.Ops(status, msg)
		if err != nil {
			return nil, err
		}

		if sdk.MsgTypeURL(msg) == MsgSendOperation {
			txOps = util.Map(txOps, func(op *types.Operation) *types.Operation {
				op.Type = TransferOperation
				return op
			})
		}
	}

	if status != StatusTxSuccess {
		return txOps, nil
	}

	// parse some specific staking related messages in order to add related operations
	// treat to other messages as generic, only process balance changing events
	var ops []*types.Operation
	switch msg := msg.(type) {
	case *distributiontypes.MsgWithdrawDelegatorReward:
		ops, err = c.processWithdrawDelegatorReward(msg, events)
	case *distributiontypes.MsgWithdrawValidatorCommission:
		ops, err = c.processWithdrawValidatorCommission(msg, events)
	case *staking.MsgDelegate:
		ops, err = c.processDelegate(msg, events)
	case *staking.MsgUndelegate:
		ops, err = c.processUndelegate(msg, events)
	case *staking.MsgBeginRedelegate:
		ops, err = c.processBeginRedelegate(msg, events)
	default:
		ops, err = c.processBalanceChange(events)
	}

	return append(txOps, ops...), err
}

func (c converter) ProcessEndBlockerEvents(events []abci.Event) []*types.Operation {
	var ops []*types.Operation

	// find complete unbonding events
	completeUnbondingEvents := util.Filter(
		events,
		func(event abci.Event) bool {
			return event.Type == staking.EventTypeCompleteUnbonding &&
				len(event.Attributes) == 3 &&
				string(event.Attributes[0].Key) == sdk.AttributeKeyAmount &&
				string(event.Attributes[1].Key) == staking.AttributeKeyValidator &&
				string(event.Attributes[2].Key) == staking.AttributeKeyDelegator
		},
	)

	delegators := util.Map(completeUnbondingEvents, func(event abci.Event) string { return string(event.Attributes[2].Value) })

	for _, e := range events {
		balanceOps, ok := c.sdkEventToBalanceOperations(StatusTxSuccess, e)
		if !ok {
			continue
		}

		balanceOps = util.Map(balanceOps, func(op *types.Operation) *types.Operation {
			if len(completeUnbondingEvents) == 0 {
				return op
			}

			if op.Account.Address == NotBondedPool.String() {
				op.Type = staking.EventTypeCompleteUnbonding
			}

			if util.Contains(delegators, op.Account.Address) {
				op.Type = staking.EventTypeCompleteUnbonding
			}

			return op
		})
		ops = append(ops, balanceOps...)
	}

	return ops
}

func (c converter) processDelegate(msg sdk.Msg, events sdk.StringEvents) ([]*types.Operation, error) {
	msgDelegate, ok := msg.(*staking.MsgDelegate)
	if !ok {
		return nil, fmt.Errorf("unexpected message type for delegation: %T", msg)
	}

	// delegations are processed in two operations: one for the delegator and one for the validator
	status := StatusTxSuccess
	opName := staking.EventTypeDelegate

	// find the pool that gets the coin
	coinReceived := util.FirstMatch(
		events,
		func(event sdk.StringEvent) bool {
			return event.Type == banktypes.EventTypeCoinReceived
		},
	)
	if coinReceived == nil {
		return nil, fmt.Errorf("could not find coin received event")
	}

	receiver := util.FirstMatch(coinReceived.Attributes, func(attr sdk.Attribute) bool {
		return attr.Value == BondedPool.String() || attr.Value == NotBondedPool.String()
	})
	if receiver == nil {
		return nil, fmt.Errorf("could not coin received to pool")
	}

	operations := []*types.Operation{
		{
			Type:    opName,
			Status:  &status,
			Account: &types.AccountIdentifier{Address: msgDelegate.DelegatorAddress},
			Amount: &types.Amount{
				Value:    "-" + msgDelegate.Amount.Amount.String(),
				Currency: c.toCurrency(msgDelegate.Amount.GetDenom()),
			},
		},
		{
			Type:              opName,
			RelatedOperations: []*types.OperationIdentifier{},
			Status:            &status,
			Account:           &types.AccountIdentifier{Address: receiver.Value},
			Amount: &types.Amount{
				Value:    msgDelegate.Amount.Amount.String(),
				Currency: c.toCurrency(msgDelegate.Amount.GetDenom()),
			},
		},
	}

	withdrawRewardsOps, err := c.parseWithdrawRewardsEvents(events)
	if err != nil {
		return nil, err
	}

	return append(operations, withdrawRewardsOps...), nil
}

func (c converter) processUndelegate(msg sdk.Msg, events sdk.StringEvents) ([]*types.Operation, error) {
	msgUndelegate, ok := msg.(*staking.MsgUndelegate)
	if !ok {
		return nil, fmt.Errorf("unexpected message type for undelegate: %T", msg)
	}

	// transfer tokens from bonded token pool to unbonded token pool
	var operations []*types.Operation
	opName := staking.EventTypeUnbond
	status := StatusTxSuccess
	operations = []*types.Operation{
		{
			Type:    opName,
			Status:  &status,
			Account: &types.AccountIdentifier{Address: BondedPool.String()},
			Amount: &types.Amount{
				Value:    "-" + msgUndelegate.Amount.Amount.String(),
				Currency: c.toCurrency(msgUndelegate.Amount.GetDenom()),
			},
		},
		{
			Type:              opName,
			RelatedOperations: []*types.OperationIdentifier{},
			Status:            &status,
			Account:           &types.AccountIdentifier{Address: NotBondedPool.String()},
			Amount: &types.Amount{
				Value:    msgUndelegate.Amount.Amount.String(),
				Currency: c.toCurrency(msgUndelegate.Amount.GetDenom()),
			},
		},
	}

	withdrawRewardsOps, err := c.parseWithdrawRewardsEvents(events)
	if err != nil {
		return nil, err
	}

	return append(operations, withdrawRewardsOps...), nil
}

// Not used
func (c converter) processCreateValidator(msg sdk.Msg, events sdk.StringEvents) ([]*types.Operation, error) {
	msgCreateValidator, ok := msg.(*staking.MsgCreateValidator)
	if !ok {
		return nil, fmt.Errorf("unexpected message type for create validator: %T", msg)
	}

	createValidatorEvent := util.FirstMatch(events, func(e sdk.StringEvent) bool { return e.Type == staking.EventTypeCreateValidator })
	if createValidatorEvent == nil {
		return nil, fmt.Errorf("withdraw rewards event not found for delegation")
	}

	opName := staking.EventTypeCreateValidator
	status := StatusTxSuccess
	operations := []*types.Operation{
		{
			Type:    opName,
			Status:  &status,
			Account: &types.AccountIdentifier{Address: msgCreateValidator.DelegatorAddress},
			Amount: &types.Amount{
				Value:    "-" + msgCreateValidator.Value.Amount.String(),
				Currency: c.toCurrency(msgCreateValidator.Value.GetDenom()),
			},
		},
		{
			Type:              opName,
			RelatedOperations: []*types.OperationIdentifier{},
			Status:            &status,
			Account:           &types.AccountIdentifier{Address: NotBondedPool.String()}, // TODO: can it be bonded right away?
			Amount: &types.Amount{
				Value:    msgCreateValidator.Value.Amount.String(),
				Currency: c.toCurrency(msgCreateValidator.Value.GetDenom()),
			},
		},
	}

	return operations, nil
}

func (c converter) processBeginRedelegate(msg sdk.Msg, events sdk.StringEvents) ([]*types.Operation, error) {
	_, ok := msg.(*staking.MsgBeginRedelegate)
	if !ok {
		return nil, fmt.Errorf("unexpected message type for create validator: %T", msg)
	}

	// potentially transfer tokens between pools
	// abstract into coin spent and coin received
	poolBalanceChange, err := c.processBalanceChange(events)
	if err != nil {
		return nil, err
	}
	poolBalanceChange = util.Filter(poolBalanceChange, func(op *types.Operation) bool {
		return op.Account.Address == NotBondedPool.String() || op.Account.Address == BondedPool.String()
	})

	withdrawRewardsEvents, err := c.parseWithdrawRewardsEvents(events)
	if err != nil {
		return nil, err
	}

	return append(poolBalanceChange, withdrawRewardsEvents...), nil
}

func (c converter) processWithdrawDelegatorReward(msg sdk.Msg, events sdk.StringEvents) ([]*types.Operation, error) {
	_, ok := msg.(*distributiontypes.MsgWithdrawDelegatorReward)
	if !ok {
		return nil, fmt.Errorf("unexpected message type for withdraw delegator reward: %T", msg)
	}

	return c.parseWithdrawRewardsEvents(events)
}

func (c converter) processWithdrawValidatorCommission(msg sdk.Msg, events sdk.StringEvents) ([]*types.Operation, error) {
	transferEvent := util.FirstMatch(events, func(e sdk.StringEvent) bool { return e.Type == banktypes.EventTypeTransfer })
	if transferEvent == nil {
		return nil, nil
	}

	delegator := util.FirstMatch(transferEvent.Attributes, func(attr sdk.Attribute) bool { return attr.Key == banktypes.AttributeKeyRecipient })
	if delegator == nil {
		return nil, fmt.Errorf("attribute recipient not found")
	}

	amountAttribute := util.FirstMatch(transferEvent.Attributes, func(attr sdk.Attribute) bool { return attr.Key == sdk.AttributeKeyAmount })
	if amountAttribute == nil {
		return nil, fmt.Errorf("attribute amount not found")
	}

	coins, err := sdk.ParseCoinsNormalized(amountAttribute.Value)
	if err != nil {
		return nil, err
	}

	var operations []*types.Operation
	// process staking rewards
	status := StatusTxSuccess
	for _, coin := range coins {
		operations = append(operations, &types.Operation{
			Type:    distributiontypes.EventTypeWithdrawCommission,
			Status:  &status,
			Account: newAccountIdentifier(Distributor.String()),
			Amount: &types.Amount{
				Value:    "-" + coin.Amount.String(),
				Currency: c.toCurrency(coin.GetDenom()),
			},
		})

		operations = append(operations, &types.Operation{
			RelatedOperations: []*types.OperationIdentifier{},
			Type:              distributiontypes.EventTypeWithdrawCommission,
			Status:            &status,
			Account:           newAccountIdentifier(delegator.Value),
			Amount: &types.Amount{
				Value:    coin.Amount.String(),
				Currency: c.toCurrency(coin.GetDenom()),
			},
		})
	}

	return operations, nil
}

// processBalanceChange only process coin spent and coin received operations
func (c converter) processBalanceChange(events sdk.StringEvents) ([]*types.Operation, error) {
	var operations []*types.Operation

	util.ForEach(events, func(event sdk.StringEvent) {
		switch event.Type {
		case banktypes.EventTypeCoinSpent:
			for i := 0; i < len(event.Attributes); i += 2 {
				spender := sdk.MustAccAddressFromBech32(event.Attributes[i].Value)
				coins, err := sdk.ParseCoinsNormalized(event.Attributes[i+1].Value)
				if err != nil {
					panic(err)
				}
				operations = append(operations, c.toOp(coins, true, spender.String(), event.Type)...)

			}

		case banktypes.EventTypeCoinReceived:
			for i := 0; i < len(event.Attributes); i += 2 {
				receiver := sdk.MustAccAddressFromBech32(event.Attributes[i].Value)
				coins, err := sdk.ParseCoinsNormalized(event.Attributes[i+1].Value)
				if err != nil {
					panic(err)
				}

				operations = append(operations, c.toOp(coins, false, receiver.String(), event.Type)...)
			}
		}

	})

	return operations, nil
}

func (c converter) toOp(coins sdk.Coins, isSub bool, account string, eventType string) []*types.Operation {
	var operations []*types.Operation
	status := StatusTxSuccess

	for _, coin := range coins {
		value := coin.Amount.String()
		// in case the event is a subtract balance one the rewrite value with
		// the negative coin identifier
		if isSub {
			value = "-" + value
		}

		op := &types.Operation{
			Type:    eventType,
			Status:  &status,
			Account: newAccountIdentifier(account),
			Amount: &types.Amount{
				Value:    value,
				Currency: c.toCurrency(coin.GetDenom()),
			},
		}
		operations = append(operations, op)
	}
	return operations
}

func (c converter) getFeeOps(events []abci.Event) []*types.Operation {
	// find the fee collection event, which is event type tx with attributes fee and fee payer
	feeEvent := util.FirstMatch(
		events,
		func(event abci.Event) bool {
			return event.Type == sdk.EventTypeTx &&
				len(event.Attributes) == 2 &&
				string(event.Attributes[0].Key) == sdk.AttributeKeyFee &&
				string(event.Attributes[1].Key) == sdk.AttributeKeyFeePayer
		},
	)

	if feeEvent == nil {
		return nil
	}

	fees, err := sdk.ParseCoinsNormalized((string)(feeEvent.Attributes[0].Value))
	if err != nil {
		panic(err)
	}

	var operations []*types.Operation
	status := StatusTxSuccess
	payer := sdk.MustAccAddressFromBech32((string)(feeEvent.Attributes[1].Value))
	for _, fee := range fees {
		operations = append(operations,
			&types.Operation{
				Type:    FeePayerOperation,
				Status:  &status,
				Account: newAccountIdentifier(payer.String()),
				Amount: &types.Amount{
					Value:    fmt.Sprintf("-%s", fee.Amount.String()),
					Currency: c.toCurrency(fee.GetDenom()),
				},
			},
			&types.Operation{
				Type:    FeeReceiverOperation,
				Status:  &status,
				Account: newAccountIdentifier(FeeCollector.String()),
				Amount: &types.Amount{
					Value:    fee.Amount.String(),
					Currency: c.toCurrency(fee.GetDenom()),
				},
			},
		)
	}

	return operations
}

// delegator can set a withdrawal address
func (c converter) parseWithdrawRewardsEvents(events sdk.StringEvents) ([]*types.Operation, error) {
	withdrawRewardsEvent := util.FirstMatch(events, func(e sdk.StringEvent) bool { return e.Type == distributiontypes.EventTypeWithdrawRewards })
	if withdrawRewardsEvent == nil {
		return nil, nil
	}

	transferEvent := util.FirstMatch(events, func(e sdk.StringEvent) bool { return e.Type == banktypes.EventTypeTransfer })
	if transferEvent == nil {
		return nil, nil
	}

	// find the withdrawal address
	var withdrawalAddr string
	for i := 0; i < len(transferEvent.Attributes); i += 3 {
		sender := transferEvent.Attributes[i+1].Value
		if sender == Distributor.String() {
			withdrawalAddr = transferEvent.Attributes[i].Value
			break
		}
	}
	if withdrawalAddr == "" {
		return nil, nil
	}

	opName := distributiontypes.EventTypeWithdrawRewards
	status := StatusTxSuccess

	var operations []*types.Operation

	for i := 0; i < len(withdrawRewardsEvent.Attributes); i += 2 {
		amountAttribute := withdrawRewardsEvent.Attributes[i]
		coins, err := sdk.ParseCoinsNormalized(amountAttribute.Value)
		if err != nil {
			return nil, err
		}
		if coins.IsZero() {
			continue
		}

		for _, coin := range coins {
			operations = append(operations, &types.Operation{
				Type:    opName,
				Status:  &status,
				Account: newAccountIdentifier(Distributor.String()),
				Amount: &types.Amount{
					Value:    "-" + coin.Amount.String(),
					Currency: c.toCurrency(coin.GetDenom()),
				},
			})

			operations = append(operations, &types.Operation{
				RelatedOperations: []*types.OperationIdentifier{},
				Type:              opName,
				Status:            &status,
				Account:           newAccountIdentifier(withdrawalAddr),
				Amount: &types.Amount{
					Value:    coin.Amount.String(),
					Currency: c.toCurrency(coin.GetDenom()),
				},
			})
		}
	}

	return operations, nil
}

func addOperationIndexes(operations []*types.Operation) []*types.Operation {
	finalOps := make([]*types.Operation, 0, len(operations))

	var currentIndex int64
	// add indexes to msg ops
	for _, op := range operations {
		op.OperationIdentifier = &types.OperationIdentifier{
			Index: currentIndex,
		}
		if op.RelatedOperations != nil {
			op.RelatedOperations = append(op.RelatedOperations, newOpID(currentIndex-1))
		}
		finalOps = append(finalOps, op)
		currentIndex++
	}

	return finalOps
}

func newOpID(index int64) *types.OperationIdentifier {
	return &types.OperationIdentifier{Index: index}
}

func newAccountIdentifier(address string) *types.AccountIdentifier {
	return &types.AccountIdentifier{Address: address}
}
