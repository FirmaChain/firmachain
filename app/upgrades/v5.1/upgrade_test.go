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

// Validator addresses (Bech32): two old validators to migrate, and two external validators
var val1Bech32Address = v5_1.OldValidators[0]
var val2Bech32Address = v5_1.OldValidators[1]
var extVal1Bech32Address = "firmavaloper1qzjafkzlfn04th8wrfg8mgrwr0rl0y6gjsx02y"
var extVal0Bech32Address = "firmavaloper1r6e67wyxms0qqgsvf2p906cexf23dl5dtekjxr"

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
	val1Addr := apptesting.MustVal(val1Bech32Address)
	val2Addr := apptesting.MustVal(val2Bech32Address)

	self1Addr := sdk.AccAddress(val1Addr)
	self2Addr := sdk.AccAddress(val1Addr)

	// External validators
	extVal0 := apptesting.MustVal(extVal0Bech32Address)
	extVal1 := apptesting.MustVal(extVal1Bech32Address)

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
	assertOldSelfState(self1Addr, val1Addr, newAcc1)
	assertOtherDelegatorsPreserved(s.TestAccs[0].Address, val1Addr) // acc1
	assertOtherDelegatorsPreserved(s.TestAccs[1].Address, val1Addr) // acc2
	assertCommissionWithdrawn(val1Addr)
	assertNewDelegation(newAcc1, newVal1)

	// For validator 2
	assertOldSelfState(self2Addr, val2Addr, newAcc2)
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

	// Create all the old validators.
	// Fund the validators standard accounts so that they can self-delegate.
	// Require also that external validators are not present inside the old validators array.
	for i := 0; i < len(v5_1.OldValidators); i++ {
		valAddrBech32 := v5_1.OldValidators[i]
		s.Require().True(extVal0Bech32Address != valAddrBech32)
		s.Require().True(extVal1Bech32Address != valAddrBech32)
		valAddr, err := sdk.ValAddressFromBech32(valAddrBech32)
		s.Require().NoError(err)
		selfAcc := sdk.AccAddress(valAddr)
		apptesting.MustFundAccount(app, ctx, selfAcc, bondDenom, 2_000_000_000)
		apptesting.MustMakeValidator(app, ctx, valAddr)
	}

	// Create all the new validators.
	// Fund the validators standard accounts so that they can self-delegate.
	// Require also that external validators are not present inside the new validators array.
	for i := 0; i < len(v5_1.Moves); i++ {
		valAddrBech32 := v5_1.Moves[i].NewValStr
		s.Require().True(extVal0Bech32Address != valAddrBech32)
		s.Require().True(extVal1Bech32Address != valAddrBech32)
		valAddr, err := sdk.ValAddressFromBech32(valAddrBech32)
		s.Require().NoError(err)
		selfAcc := sdk.AccAddress(valAddr)
		s.Require().True(selfAcc.String() == v5_1.Moves[i].NewAccStr)
		apptesting.MustFundAccount(app, ctx, selfAcc, bondDenom, 2_000_000_000)
		_, err = apptesting.MakeValidator(app, ctx, valAddr)
		if err != nil && err.Error() != fmt.Errorf("validator already exists: %s", valAddr.String()).Error() {
			panic(err)
		}
	}

	// Create the external validators
	extVal0Addr := apptesting.MustVal(extVal0Bech32Address)
	extVal1Addr := apptesting.MustVal(extVal1Bech32Address)
	apptesting.MustMakeValidator(app, ctx, extVal0Addr)
	apptesting.MustMakeValidator(app, ctx, extVal1Addr)

	// Self account addresses
	val1Addr := apptesting.MustVal(val1Bech32Address)
	val2Addr := apptesting.MustVal(val2Bech32Address)

	// Delegators account addresses
	acc1 := s.TestAccs[0].Address
	acc2 := s.TestAccs[1].Address
	s.Require().NotZero(app.AppKeepers.BankKeeper.GetBalance(ctx, acc1, bondDenom).Amount.Int64()) // TODO: use a given min amount
	s.Require().NotZero(app.AppKeepers.BankKeeper.GetBalance(ctx, acc2, bondDenom).Amount.Int64()) // TODO: use a given min amount

	// ==== Validators Setup ====
	val1 := apptesting.MustExistValidator(app, ctx, val1Addr)
	val2 := apptesting.MustExistValidator(app, ctx, val2Addr)
	validatorsToSet := []stakingtypes.Validator{val1, val2}

	for i := 0; i < len(validatorsToSet); i++ {
		val := validatorsToSet[i]
		valAddrString := val.OperatorAddress
		valAddr, err := sdk.ValAddressFromBech32(valAddrString)
		s.Require().NoError(err)
		selfAcc := sdk.AccAddress(valAddr)

		val.Commission.CommissionRates = stakingtypes.NewCommissionRates(
			math.LegacyMustNewDecFromStr("0.12"),
			math.LegacyMustNewDecFromStr("0.15"),
			math.LegacyMustNewDecFromStr("0.01"),
		)
		err = app.AppKeepers.StakingKeeper.SetValidator(ctx, val)
		s.Require().NoError(err)

		// ---- Delegations ----
		// Setup delegations for val1: self, acc1, acc2
		apptesting.MustCreateDelegation(app, ctx, selfAcc, valAddr, 1_000_000_000)
		apptesting.MustCreateDelegation(app, ctx, acc1, valAddr, 300_000_000)
		apptesting.MustCreateDelegation(app, ctx, acc2, valAddr, 400_000_000)

		// ---- Redelegations ----
		// from self to ext: this(selfAcc) -> extVal0, extVal1
		apptesting.MustCreateRedelegation(app, ctx, selfAcc, valAddr, extVal0Addr, 50_000_000)
		apptesting.MustCreateRedelegation(app, ctx, selfAcc, valAddr, extVal1Addr, 50_000_000)
		// from acc1 to ext: acc1 -> extVal0, extVal1
		apptesting.MustCreateRedelegation(app, ctx, acc1, valAddr, extVal0Addr, 20_000_000)
		apptesting.MustCreateRedelegation(app, ctx, acc1, valAddr, extVal1Addr, 20_000_000)
		// from acc2 to this: extVal0 -> val1
		// ensure acc2 has delegation at extVal0 to allow redelegation
		apptesting.MustCreateDelegation(app, ctx, acc2, extVal0Addr, 30_000_000)
		apptesting.MustCreateRedelegation(app, ctx, acc2, extVal0Addr, valAddr, 30_000_000)

		// ---- Unbonding delegations ----
		// Set up unbonding delegations from val1
		// From self: [u1,u2,u3,u4,u5,u6,u7]
		for i := 1; i <= 7; i++ {
			apptesting.MustCreateUnbondingDelegation(app, ctx, selfAcc, valAddr, math.NewInt(10_000_000))
		}
		// From acc1: [1,2,3,4,5,6,7]
		for i := 1; i <= 7; i++ {
			apptesting.MustCreateUnbondingDelegation(app, ctx, acc1, valAddr, math.NewInt(10_000_000))
		}
		// From acc2: [1,2]
		for i := 1; i <= 2; i++ {
			apptesting.MustCreateUnbondingDelegation(app, ctx, acc2, valAddr, math.NewInt(20_000_000))
		}

	}
}
