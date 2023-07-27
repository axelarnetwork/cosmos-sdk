package rosetta

import (
	"crypto/sha256"

	crgtypes "github.com/cosmos/cosmos-sdk/server/rosetta/lib/types"
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// statuses
const (
	StatusTxSuccess   = "Success"
	StatusTxReverted  = "Reverted"
	StatusPeerSynced  = "synced"
	StatusPeerSyncing = "syncing"
)

// In rosetta all state transitions must be represented as transactions
// since in tendermint begin block and end block are state transitions
// which are not represented as transactions we mock only the balance changes
// happening at those levels as transactions. (check BeginBlockTxHash for more info)
const (
	DeliverTxSize       = sha256.Size
	BeginEndBlockTxSize = DeliverTxSize + 1
	EndBlockHashStart   = 0x0
	BeginBlockHashStart = 0x1
)

// TransactionType is used to distinguish if a rosetta provided hash
// represents endblock, beginblock or deliver tx
type TransactionType int

const (
	UnrecognizedTx TransactionType = iota
	BeginBlockTx
	EndBlockTx
	DeliverTxTx
)

// metadata options

// misc
const (
	Log = "log"
)

const (
	FeePayerOperation    = "fee_payer"
	FeeReceiverOperation = "fee_receiver"
	// TransferOperation is MsgSend op
	TransferOperation = "Transfer"
	MsgSendOperation  = "/cosmos.bank.v1beta1.MsgSend"
)

// StakingOperations maps staking related operations proto name to op name
var StakingOperations = []string{
	"/cosmos.staking.v1beta1.MsgDelegate",
	"/cosmos.staking.v1beta1.MsgUndelegate",
	"/cosmos.staking.v1beta1.MsgBeginDelegate",
	"/cosmos.staking.v1beta1.MsgBeginRedelegate",
	"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
	"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission",
	distr.EventTypeWithdrawRewards,
	distr.EventTypeWithdrawCommission,
	staking.EventTypeDelegate,
	staking.EventTypeUnbond,
	staking.EventTypeRedelegate,
	staking.EventTypeCompleteUnbonding,
}

var (
	FeeCollector   = sdk.AccAddress(crypto.AddressHash([]byte(auth.FeeCollectorName)))
	Distributor    = sdk.AccAddress(crypto.AddressHash([]byte(distr.ModuleName)))
	StakingAccount = sdk.AccAddress(crypto.AddressHash([]byte(staking.ModuleName)))
	BondedPool     = sdk.AccAddress(crypto.AddressHash([]byte(staking.BondedPoolName)))
	NotBondedPool  = sdk.AccAddress(crypto.AddressHash([]byte(staking.NotBondedPoolName)))
)

// ConstructionPreprocessMetadata is used to represent
// the metadata rosetta can provide during preprocess options
type ConstructionPreprocessMetadata struct {
	Memo     string `json:"memo"`
	GasLimit uint64 `json:"gas_limit"`
	GasPrice string `json:"gas_price"`
}

func (c *ConstructionPreprocessMetadata) FromMetadata(meta map[string]interface{}) error {
	return unmarshalMetadata(meta, c)
}

// PreprocessOperationsOptionsResponse is the structured metadata options returned by the preprocess operations endpoint
type PreprocessOperationsOptionsResponse struct {
	ExpectedSigners []string `json:"expected_signers"`
	Memo            string   `json:"memo"`
	GasLimit        uint64   `json:"gas_limit"`
	GasPrice        string   `json:"gas_price"`
}

func (c PreprocessOperationsOptionsResponse) ToMetadata() (map[string]interface{}, error) {
	return marshalMetadata(c)
}

func (c *PreprocessOperationsOptionsResponse) FromMetadata(meta map[string]interface{}) error {
	return unmarshalMetadata(meta, c)
}

// ConstructionMetadata are the metadata options used to
// construct a transaction. It is returned by ConstructionMetadataFromOptions
// and fed to ConstructionPayload to process the bytes to sign.
type ConstructionMetadata struct {
	ChainID     string                 `json:"chain_id"`
	SignersData []*crgtypes.SignerData `json:"signer_data"`
	GasLimit    uint64                 `json:"gas_limit"`
	GasPrice    string                 `json:"gas_price"`
	Memo        string                 `json:"memo"`
}

func (c ConstructionMetadata) ToMetadata() (map[string]interface{}, error) {
	return marshalMetadata(c)
}

func (c *ConstructionMetadata) FromMetadata(meta map[string]interface{}) error {
	return unmarshalMetadata(meta, c)
}

// TxMetadata contains transaction memo
type TxMetadata struct {
	Memo string `json:"memo"`
}

func (c TxMetadata) ToMetadata() (map[string]interface{}, error) {
	return marshalMetadata(c)
}
