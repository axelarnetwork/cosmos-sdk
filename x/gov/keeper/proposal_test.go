package keeper_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta2"
)

func (suite *KeeperTestSuite) TestGetSetProposal() {
	tp := TestProposal
	proposal, err := suite.app.GovKeeper.SubmitProposal(suite.ctx, tp)
	suite.Require().NoError(err)
	proposalID := proposal.ProposalId
	suite.app.GovKeeper.SetProposal(suite.ctx, proposal)

	gotProposal, ok := suite.app.GovKeeper.GetProposal(suite.ctx, proposalID)
	suite.Require().True(ok)
	suite.Require().Equal(proposal, gotProposal)
}

func (suite *KeeperTestSuite) TestActivateVotingPeriod() {
	tp := TestProposal
	proposal, err := suite.app.GovKeeper.SubmitProposal(suite.ctx, tp)
	suite.Require().NoError(err)

	suite.Require().Nil(proposal.VotingStartTime)

	suite.app.GovKeeper.ActivateVotingPeriod(suite.ctx, proposal)

	proposal, ok := suite.app.GovKeeper.GetProposal(suite.ctx, proposal.ProposalId)
	suite.Require().True(ok)
	suite.Require().True(proposal.VotingStartTime.Equal(suite.ctx.BlockHeader().Time))

	activeIterator := suite.app.GovKeeper.ActiveProposalQueueIterator(suite.ctx, *proposal.VotingEndTime)
	suite.Require().True(activeIterator.Valid())

	proposalID := types.GetProposalIDFromBytes(activeIterator.Value())
	suite.Require().Equal(proposalID, proposal.ProposalId)
	activeIterator.Close()
}

type invalidProposalRoute struct{ v1beta1.TextProposal }

func (invalidProposalRoute) ProposalRoute() string { return "nonexistingroute" }

func (suite *KeeperTestSuite) TestSubmitProposal() {
	govAcct := suite.app.GovKeeper.GetGovernanceAccount(suite.ctx).GetAddress().String()
	_, _, randomAddr := testdata.KeyTestPubAddr()

	testCases := []struct {
		content     v1beta1.Content
		authority   string
		expectedErr error
	}{
		{&v1beta1.TextProposal{Title: "title", Description: "description"}, govAcct, nil},
		// Keeper does not check the validity of title and description, no error
		{&v1beta1.TextProposal{Title: "", Description: "description"}, govAcct, nil},
		{&v1beta1.TextProposal{Title: strings.Repeat("1234567890", 100), Description: "description"}, govAcct, nil},
		{&v1beta1.TextProposal{Title: "title", Description: ""}, govAcct, nil},
		{&v1beta1.TextProposal{Title: "title", Description: strings.Repeat("1234567890", 1000)}, govAcct, nil},
		// error when signer is not gov acct
		{&v1beta1.TextProposal{Title: "title", Description: "description"}, randomAddr.String(), types.ErrInvalidSigner},
		// error only when invalid route
		{&invalidProposalRoute{}, govAcct, types.ErrNoProposalHandlerExists},
	}

	for i, tc := range testCases {
		prop, err := v1beta2.NewLegacyContent(tc.content, tc.authority)
		suite.Require().NoError(err)
		_, err = suite.app.GovKeeper.SubmitProposal(suite.ctx, []sdk.Msg{prop})
		suite.Require().True(errors.Is(tc.expectedErr, err), "tc #%d; got: %v, expected: %v", i, err, tc.expectedErr)
	}
}

func (suite *KeeperTestSuite) TestGetProposalsFiltered() {
	proposalID := uint64(1)
	status := []v1beta2.ProposalStatus{v1beta2.StatusDepositPeriod, v1beta2.StatusVotingPeriod}

	addr1 := sdk.AccAddress("foo_________________")

	for _, s := range status {
		for i := 0; i < 50; i++ {
			p, err := v1beta2.NewProposal(TestProposal, proposalID, time.Now(), time.Now())
			suite.Require().NoError(err)

			p.Status = s

			if i%2 == 0 {
				d := v1beta2.NewDeposit(proposalID, addr1, nil)
				v := v1beta2.NewVote(proposalID, addr1, v1beta2.NewNonSplitVoteOption(v1beta2.OptionYes))
				suite.app.GovKeeper.SetDeposit(suite.ctx, d)
				suite.app.GovKeeper.SetVote(suite.ctx, v)
			}

			suite.app.GovKeeper.SetProposal(suite.ctx, p)
			proposalID++
		}
	}

	testCases := []struct {
		params             v1beta2.QueryProposalsParams
		expectedNumResults int
	}{
		{v1beta2.NewQueryProposalsParams(1, 50, v1beta2.StatusNil, nil, nil), 50},
		{v1beta2.NewQueryProposalsParams(1, 50, v1beta2.StatusDepositPeriod, nil, nil), 50},
		{v1beta2.NewQueryProposalsParams(1, 50, v1beta2.StatusVotingPeriod, nil, nil), 50},
		{v1beta2.NewQueryProposalsParams(1, 25, v1beta2.StatusNil, nil, nil), 25},
		{v1beta2.NewQueryProposalsParams(2, 25, v1beta2.StatusNil, nil, nil), 25},
		{v1beta2.NewQueryProposalsParams(1, 50, v1beta2.StatusRejected, nil, nil), 0},
		{v1beta2.NewQueryProposalsParams(1, 50, v1beta2.StatusNil, addr1, nil), 50},
		{v1beta2.NewQueryProposalsParams(1, 50, v1beta2.StatusNil, nil, addr1), 50},
		{v1beta2.NewQueryProposalsParams(1, 50, v1beta2.StatusNil, addr1, addr1), 50},
		{v1beta2.NewQueryProposalsParams(1, 50, v1beta2.StatusDepositPeriod, addr1, addr1), 25},
		{v1beta2.NewQueryProposalsParams(1, 50, v1beta2.StatusDepositPeriod, nil, nil), 50},
		{v1beta2.NewQueryProposalsParams(1, 50, v1beta2.StatusVotingPeriod, nil, nil), 50},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Test Case %d", i), func() {
			proposals := suite.app.GovKeeper.GetProposalsFiltered(suite.ctx, tc.params)
			suite.Require().Len(proposals, tc.expectedNumResults)

			for _, p := range proposals {
				if v1beta2.ValidProposalStatus(tc.params.ProposalStatus) {
					suite.Require().Equal(tc.params.ProposalStatus, p.Status)
				}
			}
		})
	}
}

func TestMigrateProposalMessages(t *testing.T) {
	content := v1beta1.NewTextProposal("Test", "description")
	contentMsg, err := v1beta2.NewLegacyContent(content, sdk.AccAddress("test1").String())
	require.NoError(t, err)
	content, err = v1beta2.LegacyContentFromMessage(contentMsg)
	require.NoError(t, err)
	require.Equal(t, "Test", content.GetTitle())
	require.Equal(t, "description", content.GetDescription())
}
