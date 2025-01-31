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

func postUpgradeChecks(_ *UpgradeTestSuite) {
}
