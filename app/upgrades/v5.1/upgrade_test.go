package v5_1_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	coreheader "cosmossdk.io/core/header"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/firmachain/firmachain/app/apptesting"
	appparams "github.com/firmachain/firmachain/app/params"
	v5_1 "github.com/firmachain/firmachain/app/upgrades/v5.1"
)

type UpgradeTestSuite struct {
	apptesting.TestSuite
}

func (s *UpgradeTestSuite) Setup() {
	fmt.Print("[Setup] called.\n")
}

// Setup sets up basic environment for an upgrade)
func (s *UpgradeTestSuite) SetupTest() {

	fmt.Print("[SetupTest] called.\n")

	t := s.T()
	s.Logger = log.NewLogger(os.Stderr)

	s.ChainId = "colosseum-1"
	s.BondDenom = appparams.DefaultBondDenom

	s.App, s.Ctx, s.TestAccs = apptesting.SetupApp(t, s.ChainId, s.BondDenom)
	s.StoreAccessSanityCheck()

	s.App.Commit()
	s.Ctx = s.Ctx.WithBlockHeight(1)
	s.Ctx = s.Ctx.WithHeaderInfo(coreheader.Info{Height: 1, Time: s.Ctx.BlockTime().Add(time.Second)}).WithBlockHeight(1)
	s.App.PreBlocker(s.Ctx, nil)

	s.setupValidatorState()
	s.App.Commit()

	// we do not finalize and commit, because pre-upgrade must run first.
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (s *UpgradeTestSuite) TestUpgrade() {
	preUpgradeChecks(s)

	upgradeHeight := int64(5)
	s.ConfirmUpgradeSucceeded(v5_1.UpgradeName, upgradeHeight)

	postUpgradeChecks(s)
}

func preUpgradeChecks(s *UpgradeTestSuite) {
	// Verify validators exist
	val1Addr, err := sdk.ValAddressFromBech32("firmavaloper1p90hu6pqd57xgdauf0l8dwpau73jkk7273y6ck")
	s.Require().NoError(err)
	// TODO: second one firmavaloper1x0lqg5vcynse3r6mug8vteu77cqyaqkgw2rg3x

	_, err = s.App.AppKeepers.StakingKeeper.GetValidator(s.Ctx, val1Addr)
	s.Require().NoError(err)

	s.Logger.Info("Pre-upgrade checks passed")
}

func postUpgradeChecks(s *UpgradeTestSuite) {
	ctx := s.Ctx
	k := &s.App.AppKeepers

	// Old validators and their operator (self) accounts
	val1Addr := apptesting.MustVal("firmavaloper1p90hu6pqd57xgdauf0l8dwpau73jkk7273y6ck")
	val2Addr := apptesting.MustVal("firmavaloper1x0lqg5vcynse3r6mug8vteu77cqyaqkgw2rg3x")
	self1 := apptesting.MustAcc("firma1p90hu6pqd57xgdauf0l8dwpau73jkk72qz0pcc")
	self2 := apptesting.MustAcc("firma1x0lqg5vcynse3r6mug8vteu77cqyaqkgsegn3g")

	// External validators
	extVal0 := apptesting.MustVal("firmavaloper1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq0qqqq")
	extVal1 := apptesting.MustVal("firmavaloper1wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww0wwww")

	// New accounts/validators per moves in the upgrade handler
	newAcc1 := apptesting.MustAcc("firma1k0m54qycp4v04wazj0f72snp86htau5g6ujfau")
	newVal1 := apptesting.MustVal("firmavaloper1k0m54qycp4v04wazj0f72snp86htau5gy0ejaj")
	newAcc2 := apptesting.MustAcc("firma1fzupf3r5gk505ddt4qpms0gsa7e09j68kzexzl")
	newVal2 := apptesting.MustVal("firmavaloper1fzupf3r5gk505ddt4qpms0gsa7e09j68g3jaz3")

	// Helper to assert old self state cleaned and funds moved
	assertOldSelfState := func(self sdk.AccAddress, oldVal sdk.ValAddress, newAcc sdk.AccAddress) {
		// No self delegations to old validator
		if _, err := k.StakingKeeper.GetDelegation(ctx, self, oldVal); err == nil {
			s.T().Fatal("expected no self delegation to old validator after upgrade")
		}
		// No unbonding delegations remain for self
		ubd, err := k.StakingKeeper.GetAllUnbondingDelegations(ctx, self)
		s.Require().NoError(err)
		s.Require().Equal(0, len(ubd))
		// No redelegations remain for self
		reds, err := k.StakingKeeper.GetRedelegations(ctx, self, 65535)
		s.Require().NoError(err)
		s.Require().Equal(0, len(reds))
		// Old operator account liquid should be zero (migrated)
		bal := k.BankKeeper.GetAllBalances(ctx, self)
		s.Require().True(bal.IsZero())
		// New account should hold at least the 10 ufct reserve (code keeps 10 ufct liquid)
		nb := k.BankKeeper.GetAllBalances(ctx, newAcc)
		// TODO: require it is 10 ufct
		s.Require().False(nb.IsZero())
	}

	// Helper to assert acc1/acc2 states preserved
	assertOtherDelegatorsPreserved := func(acc sdk.AccAddress, oldVal sdk.ValAddress) {
		// Delegation to old validator is still present
		del, err := k.StakingKeeper.GetDelegation(ctx, acc, oldVal)
		s.Require().NoError(err)
		s.Require().NotNil(del)
		// Unbonding entries still present for acc1 (7) or acc2 (2). We assert >0 generically here.
		ubd, err := k.StakingKeeper.GetAllUnbondingDelegations(ctx, acc)
		s.Require().NoError(err)
		s.Require().True(len(ubd) > 0)
	}

	// Assert commissions withdrawn: outstanding commissions should be zero for old validators
	assertCommissionWithdrawn := func(val sdk.ValAddress) {
		rewards, err := k.DistrKeeper.GetValidatorOutstandingRewardsCoins(ctx, val)
		s.Require().NoError(err)
		// If zero, commissions were withdrawn
		s.Require().True(rewards.IsZero())
	}

	// TODO: remove?
	// Assert new accounts delegate to new validators (upgrade code re-delegates most and keeps 10 ufct)
	assertNewDelegation := func(newAcc sdk.AccAddress, newVal sdk.ValAddress) {
		del, err := k.StakingKeeper.GetDelegation(ctx, newAcc, newVal)
		s.Require().NoError(err)
		s.Require().NotNil(del)
	}

	// For validator 1
	assertOldSelfState(self1, val1Addr, newAcc1)
	assertOtherDelegatorsPreserved(s.TestAccs[0].Address, val1Addr) // acc1
	assertOtherDelegatorsPreserved(s.TestAccs[1].Address, val1Addr) // acc2
	assertCommissionWithdrawn(val1Addr)
	assertNewDelegation(newAcc1, newVal1)

	// For validator 2
	assertOldSelfState(self2, val2Addr, newAcc2)
	assertOtherDelegatorsPreserved(s.TestAccs[0].Address, val2Addr)
	assertOtherDelegatorsPreserved(s.TestAccs[1].Address, val2Addr)
	assertCommissionWithdrawn(val2Addr)
	assertNewDelegation(newAcc2, newVal2)

	// Sanity: external validators exist (needed by pre-state) so checks above are meaningful
	_, err := k.StakingKeeper.GetValidator(ctx, extVal0)
	s.Require().NoError(err)
	_, err = k.StakingKeeper.GetValidator(ctx, extVal1)
	s.Require().NoError(err)
}

func (s *UpgradeTestSuite) setupValidatorState() {

	app := s.App
	ctx := s.Ctx
	bondDenom := s.BondDenom

	// Validator addresses (Bech32): two old validators to migrate, and two external validators
	val1Bech32Address := "firmavaloper1p90hu6pqd57xgdauf0l8dwpau73jkk7273y6ck"
	val2Bech32Address := "firmavaloper1x0lqg5vcynse3r6mug8vteu77cqyaqkgw2rg3x"
	extVal0Bech32Address := "firmavaloper1r6e67wyxms0qqgsvf2p906cexf23dl5dtekjxr"
	extVal1Bech32Address := "firmavaloper1qzjafkzlfn04th8wrfg8mgrwr0rl0y6gjsx02y"

	// Bech32 -> Validator addresses
	val1Addr, err := sdk.ValAddressFromBech32(val1Bech32Address)
	s.Require().NoError(err)
	val2Addr, err := sdk.ValAddressFromBech32(val2Bech32Address)
	s.Require().NoError(err)
	extVal0Addr, err := sdk.ValAddressFromBech32(extVal0Bech32Address)
	s.Require().NoError(err)
	extVal1Addr, err := sdk.ValAddressFromBech32(extVal1Bech32Address)
	s.Require().NoError(err)

	// Operator addresses
	selfAcc := sdk.AccAddress(val1Addr)
	selfAcc2 := sdk.AccAddress(val2Addr)
	_ = selfAcc2 // TODO: use it

	// Account addresses
	acc1 := s.TestAccs[0].Address
	acc2 := s.TestAccs[1].Address
	s.Require().NotZero(app.AppKeepers.BankKeeper.GetBalance(ctx, acc1, bondDenom).Amount.Int64()) // TODO: use a given min amount
	s.Require().NotZero(app.AppKeepers.BankKeeper.GetBalance(ctx, acc2, bondDenom).Amount.Int64()) // TODO: use a given min amount

	// Fund accounts: val1 and val2 operator accounts so they can self-delegate and redelegate
	apptesting.MustFundAccount(app, ctx, sdk.AccAddress(val1Addr), bondDenom, 2_000_000_000)
	apptesting.MustFundAccount(app, ctx, sdk.AccAddress(val2Addr), bondDenom, 2_000_000_000)

	// ==== Create Validators ====
	// Create validators: val1, val2, and two others with minimal power so that delegations/redelegations are valid
	val1 := apptesting.MustMakeValidator(app, ctx, val1Addr)
	val2 := apptesting.MustMakeValidator(app, ctx, val2Addr)
	extVal0 := apptesting.MustMakeValidator(app, ctx, extVal0Addr)
	extVal1 := apptesting.MustMakeValidator(app, ctx, extVal1Addr)
	_ = extVal0 // TODO: use it
	_ = extVal1 // TODO: use it
	// Set up commissions for val1 and val2
	val1.Commission.CommissionRates = stakingtypes.NewCommissionRates(
		math.LegacyMustNewDecFromStr("0.12"),
		math.LegacyMustNewDecFromStr("0.15"),
		math.LegacyMustNewDecFromStr("0.01"),
	)
	err = app.AppKeepers.StakingKeeper.SetValidator(ctx, val1)
	if err != nil {
		panic(err)
	}
	val2.Commission.CommissionRates = stakingtypes.NewCommissionRates(
		math.LegacyMustNewDecFromStr("0.12"),
		math.LegacyMustNewDecFromStr("0.15"),
		math.LegacyMustNewDecFromStr("0.01"),
	)
	err = app.AppKeepers.StakingKeeper.SetValidator(ctx, val2)
	if err != nil {
		panic(err)
	}

	// ==== Validator 1 setup ====

	// ---- Delegations ----
	// Setup delegations for val1: self, acc1, acc2
	apptesting.MustCreateDelegation(app, ctx, selfAcc, val1Addr, 1_000_000_000)
	apptesting.MustCreateDelegation(app, ctx, acc1, val1Addr, 100_000_000)
	apptesting.MustCreateDelegation(app, ctx, acc2, val1Addr, 200_000_000)

	// ---- Redelegations ----
	// from self to ext: this(selfAcc) -> extVal0, extVal1
	apptesting.MustCreateRedelegation(app, ctx, selfAcc, val1Addr, extVal0Addr, 50_000_000)
	apptesting.MustCreateRedelegation(app, ctx, selfAcc, val1Addr, extVal1Addr, 50_000_000)
	// from acc1 to ext: acc1 -> extVal0, extVal1
	apptesting.MustCreateRedelegation(app, ctx, acc1, val1Addr, extVal0Addr, 20_000_000)
	apptesting.MustCreateRedelegation(app, ctx, acc1, val1Addr, extVal1Addr, 20_000_000)
	// from acc2 to this: extVal0 -> val1
	// ensure acc2 has delegation at extVal0 to allow redelegation
	apptesting.MustCreateDelegation(app, ctx, acc2, extVal0Addr, 30_000_000)
	apptesting.MustCreateRedelegation(app, ctx, acc2, extVal0Addr, val1Addr, 30_000_000)

	// ---- Unbonding delegations ----
	// Set up unbonding delegations from val1
	// From self: [u1,u2,u3,u4,u5,u6,u7]
	for i := 1; i <= 7; i++ {
		apptesting.MustCreateUnbondingDelegation(app, ctx, selfAcc, val1Addr, math.NewInt(10_000_000))
	}
	// From acc1: [1,2,3,4,5,6,7]
	for i := 1; i <= 7; i++ {
		apptesting.MustCreateUnbondingDelegation(app, ctx, acc1, val1Addr, math.NewInt(10_000_000))
	}
	// From acc2: [1,2]
	for i := 1; i <= 2; i++ {
		apptesting.MustCreateUnbondingDelegation(app, ctx, acc2, val1Addr, math.NewInt(20_000_000))
	}

	// ==== Validator 2 setup ====
	// Mirror the same pre-state of val1

	// ---- Delegations ----
	apptesting.MustCreateDelegation(app, ctx, sdk.AccAddress(val2Addr), val2Addr, 1_000_000_000)
	apptesting.MustCreateDelegation(app, ctx, acc1, val2Addr, 100_000_000)
	apptesting.MustCreateDelegation(app, ctx, acc2, val2Addr, 200_000_000)

	// ---- Redelegations ----
	apptesting.MustCreateRedelegation(app, ctx, sdk.AccAddress(val2Addr), val2Addr, extVal0Addr, 50_000_000)
	apptesting.MustCreateRedelegation(app, ctx, sdk.AccAddress(val2Addr), val2Addr, extVal1Addr, 50_000_000)
	apptesting.MustCreateRedelegation(app, ctx, acc1, val2Addr, extVal0Addr, 20_000_000)
	apptesting.MustCreateRedelegation(app, ctx, acc1, val2Addr, extVal1Addr, 20_000_000)
	// acc2: extVal0 -> val2
	apptesting.MustCreateDelegation(app, ctx, acc2, extVal0Addr, 30_000_000)
	apptesting.MustCreateRedelegation(app, ctx, acc2, extVal0Addr, val2Addr, 30_000_000)

	// ---- Unbonding delegations ----
	// Unbondings for val2
	for i := 1; i <= 7; i++ {
		apptesting.MustCreateUnbondingDelegation(app, ctx, sdk.AccAddress(val2Addr), val2Addr, math.NewInt(10_000_000))
		apptesting.MustCreateUnbondingDelegation(app, ctx, acc1, val2Addr, math.NewInt(10_000_000))
	}
	for i := 1; i <= 2; i++ {
		apptesting.MustCreateUnbondingDelegation(app, ctx, acc2, val2Addr, math.NewInt(20_000_000))
	}

}
