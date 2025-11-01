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
var extVal1Bech32Address = "firmavaloper1qzjafkzlfn04th8wrfg8mgrwr0rl0y6gjsx02y"
var extVal0Bech32Address = "firmavaloper1r6e67wyxms0qqgsvf2p906cexf23dl5dtekjxr"

func allValidatorsAddresses() []string {
	addresses := append(v5_1.OldValidators, extVal0Bech32Address, extVal1Bech32Address)
	for i := 0; i < len(v5_1.Moves); i++ {
		addresses = append(addresses, v5_1.Moves[i].NewValStr)
	}
	return addresses
}

func allInternalValidatorsAddresses() []string {
	addresses := v5_1.OldValidators
	for i := 0; i < len(v5_1.Moves); i++ {
		addresses = append(addresses, v5_1.Moves[i].NewValStr)
	}
	return addresses
}

type BalStakeComm struct {
	Balances    sdk.Coins
	StakeAmount int64
	Commissions sdk.Coins
}
type UpgradeTestSuite struct {
	apptesting.TestSuite

	// Map to map the validators standard account to their balances, stake, and commission, pre-upgrade
	vAccPreUpgrade map[string]BalStakeComm
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

	s.verifyValidatorsMovesInfoWellDefined()

	s.setupValidatorState()

	s.vAccPreUpgrade = make(map[string]BalStakeComm)

	// we do not finalize and commit, because pre-upgrade must run first.
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (s *UpgradeTestSuite) TestUpgrade() {
	preUpgradeChecks(s)

	s.setupPreUpgradeBalancesMap()

	upgradeHeight := int64(5)
	s.ConfirmUpgradeSucceeded(v5_1.UpgradeName, upgradeHeight)

	postUpgradeChecks(s)
}

// Check that v5_1.Moves contains all and only the accounts to move
// associated with v5_1.OldValidators
func (s *UpgradeTestSuite) verifyValidatorsMovesInfoWellDefined() {
	valsToBeMoved := []string{}
	findValInArray := func(strArr []string, s string) (bool, int) {
		for i, v := range strArr {
			if v == s {
				return true, i
			}
		}
		return false, -1
	}

	// Save the validator addresses that will be moved inside valsToBeMoved
	// Check also that newVal and newAcc point to the same validator
	for _, mv := range v5_1.Moves {
		oldVal := sdk.ValAddress(apptesting.MustAcc(mv.OldAccStr))
		found, _ := findValInArray(valsToBeMoved, oldVal.String())
		if !found {
			valsToBeMoved = append(valsToBeMoved, oldVal.String())
		}

		newVal := sdk.ValAddress(apptesting.MustAcc(mv.NewAccStr))
		s.Require().True(mv.NewValStr == newVal.String())
	}

	// Check that each oldVal inside OldValidators is contained into v5_1.Moves
	for _, valAddrStr := range v5_1.OldValidators {
		found, pos := findValInArray(valsToBeMoved, valAddrStr)
		s.Require().True(found, fmt.Sprintf("old validator %s from v5_1.OldValidators not contained into v5_1.Moves", valAddrStr))
		valsToBeMoved = append(valsToBeMoved[:pos], valsToBeMoved[pos+1:]...)
	}

	// Check that no other validator account is inside v5_1.Moves
	s.Require().Empty(valsToBeMoved, "v5_1.Moves contains entries with old validators not present inside v5_1.OldValidators")
}

func (s *UpgradeTestSuite) setupPreUpgradeBalancesMap() {
	for _, valAddrString := range allValidatorsAddresses() {
		preupMap := s.vAccPreUpgrade
		valAddr, err := sdk.ValAddressFromBech32(valAddrString)
		s.Require().NoError(err)
		bsc := BalStakeComm{}
		bsc.Balances = s.App.AppKeepers.BankKeeper.GetAllBalances(s.Ctx, sdk.AccAddress(valAddr))
		val := apptesting.MustExistValidator(s.App, s.Ctx, valAddr)
		bsc.StakeAmount = val.GetBondedTokens().Int64()
		accumCommission, err := s.App.AppKeepers.DistrKeeper.GetValidatorAccumulatedCommission(s.Ctx, valAddr)
		s.Require().NoError(err)
		commission, _ := accumCommission.Commission.TruncateDecimal()
		bsc.Commissions = commission
		preupMap[valAddrString] = bsc
		s.vAccPreUpgrade = preupMap
	}
}

func (s *UpgradeTestSuite) verifyAllValidatorsExist() {
	validatorsToCheck := allValidatorsAddresses()
	for i := 0; i < len(validatorsToCheck); i++ {
		valAddrBech32 := validatorsToCheck[i]
		valAddr, err := sdk.ValAddressFromBech32(valAddrBech32)
		s.Require().NoError(err)
		apptesting.MustExistValidator(s.App, s.Ctx, valAddr)
	}
}

func (s *UpgradeTestSuite) verifyValidatorsCantUndelegate(validatorsToCheck []string) {
	for i := 0; i < len(validatorsToCheck); i++ {
		valAddrBech32 := validatorsToCheck[i]
		valAddr, err := sdk.ValAddressFromBech32(valAddrBech32)
		s.Require().NoError(err)
		selfAddr := sdk.AccAddress(valAddr)
		_, _, err = apptesting.CreateUnbondingDelegation(s.App, s.Ctx, selfAddr, valAddr, 1)
		s.Require().Error(err, fmt.Sprintf("Validator can undelegate: %s", valAddrBech32))
	}
}

func preUpgradeChecks(s *UpgradeTestSuite) {
	s.verifyAllValidatorsExist()
	s.verifyValidatorsCantUndelegate(allInternalValidatorsAddresses())

	s.Logger.Info("Pre-upgrade checks passed")
}

func postUpgradeChecks(s *UpgradeTestSuite) {
	ctx := s.Ctx
	k := &s.App.AppKeepers

	s.verifyAllValidatorsExist()

	// Helper to assert old self state cleaned and funds moved
	// Assert that, for the account (self), there are no:
	// - self delegation
	// - unbonding delegations
	// - redelegations
	// - balances
	assertOldValState := func(oldVal sdk.ValAddress) {
		self := sdk.AccAddress(oldVal)
		if _, err := k.StakingKeeper.GetDelegation(ctx, self, oldVal); err == nil {
			s.T().Fatal("expected no self delegation to old validator after upgrade")
		}
		ubd, err := k.StakingKeeper.GetAllUnbondingDelegations(ctx, self)
		s.Require().NoError(err)
		s.Require().Equal(0, len(ubd))
		reds, err := k.StakingKeeper.GetRedelegations(ctx, self, 65535)
		s.Require().NoError(err)
		s.Require().Equal(0, len(reds))
		bal := k.BankKeeper.GetAllBalances(ctx, self)
		s.Require().True(bal.IsZero())
	}

	assertNewValState := func(newVal sdk.ValAddress) {
		self := sdk.AccAddress(newVal)
		liquidAmount := k.BankKeeper.GetBalance(ctx, self, s.BondDenom).Amount.Int64()
		// Liquid amount must not be greater than v5_1.NotDelegatedAmount
		s.Require().LessOrEqual(liquidAmount, v5_1.NotDelegatedAmount, fmt.Sprintf("new validator: %s", newVal.String()))
		del, err := k.StakingKeeper.GetDelegation(ctx, self, newVal)
		// If liquid amount is less than v5_1.NotDelegatedAmount, there should not be a new self-delegation in the new validator.
		if liquidAmount < v5_1.NotDelegatedAmount {
			s.Require().Error(err)
		} else {
			s.Require().NoError(err)
			s.Require().NotNil(del)
		}
	}

	// Helper to assert acc1/acc2 states preserved
	assertOtherDelegatorsPreserved := func(delegators []sdk.AccAddress, oldVal sdk.ValAddress) {
		for _, d := range delegators {
			// Delegation to old validator is still present
			del, err := k.StakingKeeper.GetDelegation(ctx, d, oldVal)
			s.Require().NoError(err)
			s.Require().NotNil(del)
			// Unbonding entries still present for acc1 (7) or acc2 (2). We assert >0 generically here.
			ubd, err := k.StakingKeeper.GetAllUnbondingDelegations(ctx, d)
			s.Require().NoError(err)
			s.Require().True(len(ubd) > 0)
		}
	}

	// Assert commissions withdrawn: outstanding commissions should be zero for old validators
	assertCommissionWithdrawn := func(val sdk.ValAddress) {
		rewards, err := k.DistrKeeper.GetValidatorOutstandingRewardsCoins(ctx, val)
		s.Require().NoError(err)
		// If zero, commissions were withdrawn
		s.Require().True(rewards.IsZero())
	}

	for i := 0; i < len(v5_1.Moves); i++ {
		oldAccStr := v5_1.Moves[i].OldAccStr
		newValStr := v5_1.Moves[i].NewValStr
		oldAccAddr, err := sdk.AccAddressFromBech32(oldAccStr)
		s.Require().NoError(err)
		oldValAddr := sdk.ValAddress(oldAccAddr)
		s.Require().NoError(err)
		newValAddr, err := sdk.ValAddressFromBech32(newValStr)
		s.Require().NoError(err)

		assertOldValState(oldValAddr)
		assertOtherDelegatorsPreserved([]sdk.AccAddress{s.TestAccs[0].Address, s.TestAccs[1].Address}, oldValAddr)
		// TODO: set commissions
		assertCommissionWithdrawn(oldValAddr)
		assertNewValState(newValAddr)
	}
}

func (s *UpgradeTestSuite) setupValidatorState() {

	app := s.App
	ctx := s.Ctx
	bondDenom := s.BondDenom

	// Create all the old validators.
	// Fund the validators standard accounts so that they can self-delegate.
	// Require also that external validators are not present inside the old validators array.
	oldValidators := []stakingtypes.Validator{}
	for i := 0; i < len(v5_1.OldValidators); i++ {
		valAddrBech32 := v5_1.OldValidators[i]
		s.Require().True(extVal0Bech32Address != valAddrBech32)
		s.Require().True(extVal1Bech32Address != valAddrBech32)
		valAddr, err := sdk.ValAddressFromBech32(valAddrBech32)
		s.Require().NoError(err)
		selfAcc := sdk.AccAddress(valAddr)
		apptesting.MustFundAccount(app, ctx, selfAcc, bondDenom, 2_000_000_000_000)
		v := apptesting.MustMakeValidator(app, ctx, valAddr)
		oldValidators = append(oldValidators, v)
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
		// If we already set the validator, avoid funding again the account
		_, err = app.AppKeepers.StakingKeeper.GetValidator(ctx, valAddr)
		if err == nil {
			continue
		}
		selfAcc := sdk.AccAddress(valAddr)
		s.Require().True(selfAcc.String() == v5_1.Moves[i].NewAccStr)
		apptesting.MustFundAccount(app, ctx, selfAcc, bondDenom, 2_000_000_000)
		apptesting.MustMakeValidator(app, ctx, valAddr)
	}

	// Create the external validators
	extVal0Addr := apptesting.MustVal(extVal0Bech32Address)
	extVal1Addr := apptesting.MustVal(extVal1Bech32Address)
	apptesting.MustMakeValidator(app, ctx, extVal0Addr)
	apptesting.MustMakeValidator(app, ctx, extVal1Addr)

	// Delegators account addresses
	acc1 := s.TestAccs[0].Address
	acc2 := s.TestAccs[1].Address
	// these accounts should be funded by the suite setup process
	s.Require().NotZero(app.AppKeepers.BankKeeper.GetBalance(ctx, acc1, bondDenom).Amount.Int64())
	s.Require().NotZero(app.AppKeepers.BankKeeper.GetBalance(ctx, acc2, bondDenom).Amount.Int64())

	// ==== Validators Setup ====
	maxEntries, err := s.App.AppKeepers.StakingKeeper.MaxEntries(ctx)
	s.Require().NoError(err)
	fmt.Printf("staking params: maxEntries: %v\n", maxEntries)
	for i := 0; i < len(oldValidators); i++ {
		val := oldValidators[i]
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
		for i := 1; i <= int(maxEntries); i++ {
			// dummy-increment blockHeight, otherwise entries are summed up
			ctx = s.Ctx.WithBlockHeight(s.Ctx.BlockHeight() + int64(i) - 1)
			apptesting.MustCreateUnbondingDelegation(app, ctx, selfAcc, valAddr, 10_000_000)
		}

		// From acc1: [1,2,3,4,5,6,7]
		for i := 1; i <= int(maxEntries); i++ {
			// dummy-increment blockHeight, otherwise entries are summed up
			ctx = s.Ctx.WithBlockHeight(s.Ctx.BlockHeight() + int64(i) - 1)
			apptesting.MustCreateUnbondingDelegation(app, ctx, acc1, valAddr, 10_000_000)
		}
		// From acc2: [1,2]
		for i := 1; i <= 2; i++ {
			// dummy-increment blockHeight, otherwise entries are summed up
			ctx = s.Ctx.WithBlockHeight(s.Ctx.BlockHeight() + int64(i) - 1)
			apptesting.MustCreateUnbondingDelegation(app, ctx, acc2, valAddr, 20_000_000)
		}

	}
}
