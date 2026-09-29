package v5_2_test

import (
	"os"
	"testing"

	"cosmossdk.io/log"
	"github.com/stretchr/testify/suite"

	"github.com/firmachain/firmachain/app/apptesting"
	appparams "github.com/firmachain/firmachain/app/params"
	v5_2 "github.com/firmachain/firmachain/app/upgrades/v5.2"
)

type UpgradeTestSuite struct {
	apptesting.TestSuite
}

func (s *UpgradeTestSuite) SetupTest() {
	s.Logger = log.NewLogger(os.Stderr)
	s.ChainId = "colosseum-1"
	s.BondDenom = appparams.DefaultBondDenom
	s.App, s.Ctx, s.TestAccs = apptesting.SetupApp(s.T(), s.ChainId, s.BondDenom)
	s.StoreAccessSanityCheck()
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (s *UpgradeTestSuite) TestUpgrade() {
	versionMapBefore, err := s.App.AppKeepers.UpgradeKeeper.GetModuleVersionMap(s.Ctx)
	s.Require().NoError(err)

	s.ConfirmUpgradeSucceeded(v5_2.UpgradeName, 5)

	versionMapAfter, err := s.App.AppKeepers.UpgradeKeeper.GetModuleVersionMap(s.Ctx)
	s.Require().NoError(err)
	s.Require().Equal(versionMapBefore, versionMapAfter)
}
