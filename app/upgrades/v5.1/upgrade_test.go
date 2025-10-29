package v5_1_test

import (
	"os"
	"testing"
	"time"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	coreheader "cosmossdk.io/core/header"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtypes "github.com/cometbft/cometbft/types"

	"github.com/firmachain/firmachain/app"
	"github.com/firmachain/firmachain/app/apptesting"
	apphelpers "github.com/firmachain/firmachain/app/helpers"
	appparams "github.com/firmachain/firmachain/app/params"
	v5_1 "github.com/firmachain/firmachain/app/upgrades/v5.1"
)

type UpgradeTestSuite struct {
	apptesting.TestSuite
}

func (s *UpgradeTestSuite) SetupTest() {
	s.Logger = log.NewLogger(os.Stderr)

	s.ChainId = "colosseum-1"
	s.BondDenom = appparams.DefaultBondDenom

	s.App, s.Ctx, s.TestAccs = setupAppWithValidators(s.T(), s.ChainId, s.BondDenom)
	s.StoreAccessSanityCheck()
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

	_, err = s.App.AppKeepers.StakingKeeper.GetValidator(s.Ctx, val1Addr)
	s.Require().NoError(err)

	s.Logger.Info("Pre-upgrade checks passed")
}

func postUpgradeChecks(s *UpgradeTestSuite) {
	ctx := s.Ctx
	k := &s.App.AppKeepers

	// Old validators and their operator (self) accounts
	val1Addr := mustVal("firmavaloper1p90hu6pqd57xgdauf0l8dwpau73jkk7273y6ck")
	val2Addr := mustVal("firmavaloper1x0lqg5vcynse3r6mug8vteu77cqyaqkgw2rg3x")
	self1 := mustAcc("firma1p90hu6pqd57xgdauf0l8dwpau73jkk72qz0pcc")
	self2 := mustAcc("firma1x0lqg5vcynse3r6mug8vteu77cqyaqkgsegn3g")

	// External validators
	extVal0 := mustVal("firmavaloper1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq0qqqq")
	extVal1 := mustVal("firmavaloper1wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww0wwww")

	// New accounts/validators per moves in the upgrade handler
	newAcc1 := mustAcc("firma1k0m54qycp4v04wazj0f72snp86htau5g6ujfau")
	newVal1 := mustVal("firmavaloper1k0m54qycp4v04wazj0f72snp86htau5gy0ejaj")
	newAcc2 := mustAcc("firma1fzupf3r5gk505ddt4qpms0gsa7e09j68kzexzl")
	newVal2 := mustVal("firmavaloper1fzupf3r5gk505ddt4qpms0gsa7e09j68g3jaz3")

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

// Helpers for parsing addresses safely in post-checks
func mustAcc(s string) sdk.AccAddress {
	a, err := sdk.AccAddressFromBech32(s)
	if err != nil {
		panic(err)
	}
	return a
}

func mustVal(s string) sdk.ValAddress {
	v, err := sdk.ValAddressFromBech32(s)
	if err != nil {
		panic(err)
	}
	return v
}

func setupAppWithValidators(t *testing.T, chainId string, bondDenom string) (*app.App, sdk.Context, []apptesting.AddressWithKeys) {
	appparams.SetSdkConfigAndSeal()

	// Create test accounts
	testAddrsWithKeys := apptesting.CreateRandomAccounts(5)

	// Create validator set
	privVal := apphelpers.NewPV()
	pubKey, err := privVal.GetPubKey()
	if err != nil {
		panic(err)
	}

	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// Create genesis accounts and balances
	genesisAccounts := make([]authtypes.GenesisAccount, 0, len(testAddrsWithKeys))
	genesisBalances := make([]banktypes.Balance, 0, len(testAddrsWithKeys))

	for _, acc := range testAddrsWithKeys {
		acc := authtypes.NewBaseAccount(acc.Address.Bytes(), acc.PubKey, 0, 0)
		balance := banktypes.Balance{
			Address: acc.GetAddress().String(),
			Coins:   sdk.NewCoins(sdk.NewCoin(bondDenom, math.NewInt(100000000000000))),
		}
		genesisAccounts = append(genesisAccounts, acc)
		genesisBalances = append(genesisBalances, balance)
	}

	timenow := time.Now()
	initialHeight := int64(1)

	app := apptesting.SetupWithGenesisValSet(t, chainId, bondDenom, timenow, initialHeight, valSet, genesisAccounts, genesisBalances...)

	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{
		ChainID: chainId,
		Height:  initialHeight,
		Time:    timenow,
	}).WithHeaderInfo(coreheader.Info{
		ChainID: chainId,
		Height:  initialHeight,
		Time:    timenow,
	})

	// Set up the specific validator state
	setupValidatorState(app, ctx, testAddrsWithKeys)

	return app, ctx, testAddrsWithKeys
}

func setupValidatorState(app *app.App, ctx sdk.Context, testAccs []apptesting.AddressWithKeys) {
	bondDenom := appparams.DefaultBondDenom

	// Validator addresses (two old validators to migrate) and two external validators
	val1Addr, _ := sdk.ValAddressFromBech32("firmavaloper1p90hu6pqd57xgdauf0l8dwpau73jkk7273y6ck")
	val2Addr, _ := sdk.ValAddressFromBech32("firmavaloper1x0lqg5vcynse3r6mug8vteu77cqyaqkgw2rg3x")
	extVal0Addr, _ := sdk.ValAddressFromBech32("firmavaloper1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq0qqqq")
	extVal1Addr, _ := sdk.ValAddressFromBech32("firmavaloper1wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww0wwww")

	// Account addresses
	selfAcc, _ := sdk.AccAddressFromBech32("firma1p90hu6pqd57xgdauf0l8dwpau73jkk72qz0pcc")
	acc1, _ := sdk.AccAddressFromBech32(testAccs[0].Address.String()) // Use first test account
	acc2, _ := sdk.AccAddressFromBech32(testAccs[1].Address.String()) // Use second test account

	// Ensure validators exist (val1 from genesis; create val2/ext validators for the test)
	val1, err := app.AppKeepers.StakingKeeper.GetValidator(ctx, val1Addr)
	if err != nil {
		panic(err)
	}
	// Create val2 and two external validators with minimal power so that delegations/redelegations are valid
	mustEnsureValidator(app, ctx, val2Addr)
	mustEnsureValidator(app, ctx, extVal0Addr)
	mustEnsureValidator(app, ctx, extVal1Addr)

	// Fund operator accounts (self) so they can self-delegate and redelegate
	mustFundAccount(app, ctx, sdk.AccAddress(val1Addr), bondDenom, math.NewInt(2_000_000_000))
	mustFundAccount(app, ctx, sdk.AccAddress(val2Addr), bondDenom, math.NewInt(2_000_000_000))

	// Delegate from self to val1 (1000 tokens)
	delegateAmount := math.NewInt(1000000000)
	_, err = app.AppKeepers.StakingKeeper.Delegate(ctx, selfAcc, delegateAmount, stakingtypes.Unbonded, val1, true)
	if err != nil {
		panic(err)
	}

	// Delegate from acc1 to val1 (100 tokens)
	acc1Delegation := math.NewInt(100000000)
	_, err = app.AppKeepers.StakingKeeper.Delegate(ctx, acc1, acc1Delegation, stakingtypes.Unbonded, val1, true)
	if err != nil {
		panic(err)
	}

	// Delegate from acc2 to val1 (200 tokens)
	acc2Delegation := math.NewInt(200000000)
	_, err = app.AppKeepers.StakingKeeper.Delegate(ctx, acc2, acc2Delegation, stakingtypes.Unbonded, val1, true)
	if err != nil {
		panic(err)
	}

	// Set up commissions for val1
	val1.Commission.CommissionRates = stakingtypes.NewCommissionRates(
		math.LegacyMustNewDecFromStr("0.12"),
		math.LegacyMustNewDecFromStr("0.15"),
		math.LegacyMustNewDecFromStr("0.01"),
	)
	err = app.AppKeepers.StakingKeeper.SetValidator(ctx, val1)
	if err != nil {
		panic(err)
	}

	// Set liquid balance for val1 operator
	val1OperatorAcc := sdk.AccAddress(val1Addr)
	err = app.AppKeepers.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(bondDenom, math.NewInt(123000000))))
	if err != nil {
		panic(err)
	}
	err = app.AppKeepers.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, val1OperatorAcc, sdk.NewCoins(sdk.NewCoin(bondDenom, math.NewInt(123000000))))
	if err != nil {
		panic(err)
	}

	// ==== Redelegations per schedule ====
	// from self to ext: this(selfAcc) -> extVal0, extVal1
	_, err = app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, selfAcc, val1Addr, extVal0Addr, math.LegacyNewDec(50_000_000))
	if err != nil {
		panic(err)
	}
	_, err = app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, selfAcc, val1Addr, extVal1Addr, math.LegacyNewDec(50_000_000))
	if err != nil {
		panic(err)
	}
	// from acc1 to ext: acc1 -> extVal0, extVal1
	_, err = app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, acc1, val1Addr, extVal0Addr, math.LegacyNewDec(20_000_000))
	if err != nil {
		panic(err)
	}
	_, err = app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, acc1, val1Addr, extVal1Addr, math.LegacyNewDec(20_000_000))
	if err != nil {
		panic(err)
	}
	// from acc2 to this: extVal0 -> val1
	// ensure acc2 has delegation at extVal0 to allow redelegation
	ext0, _ := app.AppKeepers.StakingKeeper.GetValidator(ctx, extVal0Addr)
	_, err = app.AppKeepers.StakingKeeper.Delegate(ctx, acc2, math.NewInt(30_000_000), stakingtypes.Unbonded, ext0, true)
	if err != nil {
		panic(err)
	}
	_, err = app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, acc2, extVal0Addr, val1Addr, math.LegacyNewDec(30_000_000))
	if err != nil {
		panic(err)
	}

	// Set up unbonding delegations from val1
	// From self: [u1,u2,u3,u4,u5,u6,u7]
	for i := 1; i <= 7; i++ {
		createUnbondingDelegation(app, ctx, selfAcc, val1Addr, math.NewInt(10000000))
	}

	// From acc1: [1,2,3,4,5,6,7]
	for i := 1; i <= 7; i++ {
		createUnbondingDelegation(app, ctx, acc1, val1Addr, math.NewInt(10000000))
	}

	// From acc2: [1,2]
	for i := 1; i <= 2; i++ {
		createUnbondingDelegation(app, ctx, acc2, val1Addr, math.NewInt(20000000))
	}

	// ==== Mirror the same pre-state for validator 2 ====
	// Delegators
	val2, _ := app.AppKeepers.StakingKeeper.GetValidator(ctx, val2Addr)
	_, err = app.AppKeepers.StakingKeeper.Delegate(ctx, sdk.AccAddress(val2Addr), math.NewInt(1_000_000_000), stakingtypes.Unbonded, val2, true)
	if err != nil {
		panic(err)
	}
	_, err = app.AppKeepers.StakingKeeper.Delegate(ctx, acc1, math.NewInt(100_000_000), stakingtypes.Unbonded, val2, true)
	if err != nil {
		panic(err)
	}
	_, err = app.AppKeepers.StakingKeeper.Delegate(ctx, acc2, math.NewInt(200_000_000), stakingtypes.Unbonded, val2, true)
	if err != nil {
		panic(err)
	}
	// Commissions and liquid balance for val2
	val2.Commission.CommissionRates = stakingtypes.NewCommissionRates(
		math.LegacyMustNewDecFromStr("0.12"),
		math.LegacyMustNewDecFromStr("0.15"),
		math.LegacyMustNewDecFromStr("0.01"),
	)
	_ = app.AppKeepers.StakingKeeper.SetValidator(ctx, val2)
	mustFundAccount(app, ctx, sdk.AccAddress(val2Addr), bondDenom, math.NewInt(123_000_000))
	// Redelegations
	_, _ = app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, sdk.AccAddress(val2Addr), val2Addr, extVal0Addr, math.LegacyNewDec(50_000_000))
	_, _ = app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, sdk.AccAddress(val2Addr), val2Addr, extVal1Addr, math.LegacyNewDec(50_000_000))
	_, _ = app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, acc1, val2Addr, extVal0Addr, math.LegacyNewDec(20_000_000))
	_, _ = app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, acc1, val2Addr, extVal1Addr, math.LegacyNewDec(20_000_000))
	// acc2: extVal0 -> val2
	_, _ = app.AppKeepers.StakingKeeper.Delegate(ctx, acc2, math.NewInt(30_000_000), stakingtypes.Unbonded, ext0, true)
	_, _ = app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, acc2, extVal0Addr, val2Addr, math.LegacyNewDec(30_000_000))
	// Unbondings for val2
	for i := 1; i <= 7; i++ {
		createUnbondingDelegation(app, ctx, sdk.AccAddress(val2Addr), val2Addr, math.NewInt(10_000_000))
		createUnbondingDelegation(app, ctx, acc1, val2Addr, math.NewInt(10_000_000))
	}
	for i := 1; i <= 2; i++ {
		createUnbondingDelegation(app, ctx, acc2, val2Addr, math.NewInt(20_000_000))
	}
}

func createUnbondingDelegation(app *app.App, ctx sdk.Context, delegator sdk.AccAddress, valAddr sdk.ValAddress, amount math.Int) {
	shares := math.LegacyNewDecFromInt(amount)
	_, _, err := app.AppKeepers.StakingKeeper.Undelegate(ctx, delegator, valAddr, shares)
	if err != nil {
		panic(err)
	}
}

// mustEnsureValidator creates a minimal validator if it does not exist yet
func mustEnsureValidator(app *app.App, ctx sdk.Context, valAddr sdk.ValAddress) {
	if _, err := app.AppKeepers.StakingKeeper.GetValidator(ctx, valAddr); err == nil {
		return
	}
	// Create a dummy validator with small power
	val := stakingtypes.Validator{
		OperatorAddress: valAddr.String(),
		Status:          stakingtypes.Bonded,
		Tokens:          math.NewInt(1),
		DelegatorShares: math.LegacyOneDec(),
		Commission:      stakingtypes.NewCommission(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()),
	}
	_ = app.AppKeepers.StakingKeeper.SetValidator(ctx, val)
}

// mustFundAccount: utility per mint + send to target account.
func mustFundAccount(app *app.App, ctx sdk.Context, addr sdk.AccAddress, denom string, amount math.Int) {
	_ = app.AppKeepers.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, amount)))
	_ = app.AppKeepers.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, sdk.NewCoins(sdk.NewCoin(denom, amount)))
}
