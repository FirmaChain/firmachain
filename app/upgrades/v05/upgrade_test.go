package v05_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/firmachain/firmachain/v05/app/apptesting"
	v05 "github.com/firmachain/firmachain/v05/app/upgrades/v05"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type UpgradeTestSuite struct {
	apptesting.TestSuite
}

func (s *UpgradeTestSuite) SetupTest() {
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

	initialVersionMap := s.App.AppKeepers.UpgradeKeeper.GetInitVersionMap()
	for k, v := range initialVersionMap {
		s.Logger.Debug(fmt.Sprintf("initialVersionMap: %s %d\n", k, v))
	}

}

func postUpgradeChecks(_ *UpgradeTestSuite) {
}
