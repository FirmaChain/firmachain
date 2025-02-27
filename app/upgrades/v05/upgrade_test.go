package v05_test

import (
	"fmt"
	"os"
	"testing"

	"cosmossdk.io/log"
	"github.com/stretchr/testify/suite"

	"github.com/firmachain/firmachain/v05/app/apptesting"
	v05 "github.com/firmachain/firmachain/v05/app/upgrades/v05"

	module "github.com/cosmos/cosmos-sdk/types/module"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

type UpgradeTestSuite struct {
	apptesting.TestSuite
}

func (s *UpgradeTestSuite) SetupTest() {
}

// Setup sets up basic environment for an upgrade)
func (s *UpgradeTestSuite) Setup() {
	t := s.T()
	s.Logger = log.NewLogger(os.Stderr)

	s.App, s.Ctx, s.TestAccs = apptesting.SetupApp(t)
	s.StoreAccessSanityCheck()

	// we do not finalize and commit, because pre-upgrade must run first.
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

// Ensures the test does not error out.
func (s *UpgradeTestSuite) TestUpgrade() {
	s.Setup()
	preUpgradeChecks(s)

	preUpgradeProposalsSetter(s)

	upgradeHeight := int64(5)
	s.ConfirmUpgradeSucceeded(v05.UpgradeName, upgradeHeight)

	postUpgradeChecks(s)
}

func preUpgradeChecks(s *UpgradeTestSuite) {
	stakingParams, err := s.App.AppKeepers.StakingKeeper.GetParams(s.Ctx)
	_ = stakingParams
	s.Logger.Debug(fmt.Sprintf("stakingParams %v", stakingParams))
	s.Require().NoError(err)

	bankParams := s.App.AppKeepers.BankKeeper.GetParams(s.Ctx)
	_ = bankParams
	s.Logger.Debug(fmt.Sprintf("bankParams %v", bankParams))

	totalSupply, err := s.App.AppKeepers.BankKeeper.TotalSupply(s.Ctx, &banktypes.QueryTotalSupplyRequest{})
	_ = totalSupply
	s.Logger.Debug(fmt.Sprintf("totalSupply %v", totalSupply))
	s.Require().NoError(err)

	balances_0, err := s.App.AppKeepers.BankKeeper.AllBalances(s.Ctx, banktypes.NewQueryAllBalancesRequest(s.TestAccs[0].Address, nil, false))
	_ = balances_0
	s.Logger.Debug(fmt.Sprintf("balances_0 %v", balances_0))
	s.Require().NoError(err)

	vm, err := s.App.AppKeepers.UpgradeKeeper.GetModuleVersionMap(s.Ctx)
	s.NoError(err)
	for v, i := range s.App.ModuleManager().Modules {
		if i, ok := i.(module.HasConsensusVersion); ok {
			s.Equal(vm[v], i.ConsensusVersion())
		}
	}
	s.NotNil(s.App.AppKeepers.UpgradeKeeper.GetVersionSetter())

}

func postUpgradeChecks(s *UpgradeTestSuite) {
	// Verify gov params have been correctly set.
	govParams, err := s.App.AppKeepers.GovKeeper.Params.Get(s.Ctx)
	s.Require().NoError(err)
	s.Require().NoError(err)
	s.Assert().Equal("0.520000000000000000", govParams.MinInitialDepositRatio)

	// Verify gov proposer field has been set correctly during the proposals migration.
	proposals := s.App.AppKeepers.GovKeeper.Proposals
	p, err := proposals.Get(s.Ctx, 1)
	s.Assert().NoError(err)
	s.Assert().Equal("firma1w02lwp4ptdk2e6gzn82jezkjznzh88a9wusumf", p.Proposer)

	p, err = proposals.Get(s.Ctx, 2)
	s.Assert().NoError(err)
	s.Assert().Equal("firma1yl76cswscpcpjcljpavqk9v5pjrtqxpy9nlrd4", p.Proposer)

	p, err = proposals.Get(s.Ctx, 3)
	s.Assert().NoError(err)
	s.Assert().Equal("firma1w02lwp4ptdk2e6gzn82jezkjznzh88a9wusumf", p.Proposer)

	p, err = proposals.Get(s.Ctx, 8)
	s.Assert().NoError(err)
	s.Assert().Equal("firma1w02lwp4ptdk2e6gzn82jezkjznzh88a9wusumf", p.Proposer)
}

// This is used to set up the pre-upgrade state for the gov proposals migtration
// (adding the proposer field via `govv4.AddProposerAddressToProposal(ctx, govKVStoreService, appCodec, proposalsIdToProposerAddress)`).
func preUpgradeProposalsSetter(s *UpgradeTestSuite) {
	dummyProposer, err := sdk.AccAddressFromBech32("firma1w02lwp4ptdk2e6gzn82jezkjznzh88a9wusumf")
	s.Require().NoError(err)

	// We register the proposals with the same id that will be used in the migration, with a dummy proposer.
	registerProposal(s, 1, dummyProposer)
	registerProposal(s, 2, dummyProposer)
	registerProposal(s, 3, dummyProposer)
	registerProposal(s, 8, dummyProposer)
}
func registerProposal(s *UpgradeTestSuite, id uint64, proposer sdk.AccAddress) {
	prop, err := govtypesv1.NewProposal(
		[]sdk.Msg{}, id, s.Ctx.BlockTime(), s.Ctx.BlockTime(), "meta", "title", "summary", proposer, false,
	)
	s.Require().NoError(err)
	err = s.App.AppKeepers.GovKeeper.SetProposal(s.Ctx, prop)
	s.Require().NoError(err)
}
