package v5_1_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/firmachain/firmachain/app/apptesting"
	appparams "github.com/firmachain/firmachain/app/params"
	v5_1 "github.com/firmachain/firmachain/app/upgrades/v5.1"
)

// ensureValidator creates a minimal validator if it does not exist.
func ensureValidator(t *testing.T, s *apptesting.TestSuite, valAddr sdk.ValAddress) stakingtypes.Validator {
	t.Helper()
	v, err := s.App.AppKeepers.StakingKeeper.GetValidator(s.Ctx, valAddr)
	if err == nil {
		return v
	}
	v = stakingtypes.Validator{
		OperatorAddress: valAddr.String(),
		Status:          stakingtypes.Bonded,
		Tokens:          math.NewInt(1),
		DelegatorShares: math.LegacyOneDec(),
		Commission:      stakingtypes.NewCommission(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()),
	}
	_ = s.App.AppKeepers.StakingKeeper.SetValidator(s.Ctx, v)
	return v
}

func fund(t *testing.T, s *apptesting.TestSuite, addr sdk.AccAddress, amt int64) {
	t.Helper()
	denom := appparams.DefaultBondDenom
	coins := sdk.NewCoins(sdk.NewCoin(denom, math.NewInt(amt)))
	_ = s.App.AppKeepers.BankKeeper.MintCoins(s.Ctx, minttypes.ModuleName, coins)
	_ = s.App.AppKeepers.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, minttypes.ModuleName, addr, coins)
}

func TestCompleteAllRedelegations(t *testing.T) {
	var s apptesting.TestSuite
	s.Setup()

	denom := appparams.DefaultBondDenom
	delegator := s.TestAccs[0].Address

	// Create validators: src and two destinations
	srcVal := sdk.ValAddress([]byte("src-validator-addr-__________1"))
	dstVal0 := sdk.ValAddress([]byte("dst-validator-addr-__________0"))
	dstVal1 := sdk.ValAddress([]byte("dst-validator-addr-__________1"))
	vSrc := ensureValidator(t, &s, srcVal)
	_ = ensureValidator(t, &s, dstVal0)
	_ = ensureValidator(t, &s, dstVal1)

	// Fund delegator and delegate to src
	fund(t, &s, delegator, 2_000_000_000)
	_, err := s.App.AppKeepers.StakingKeeper.Delegate(s.Ctx, delegator, math.NewInt(1_000_000_000), stakingtypes.Unbonded, vSrc, true)
	if err != nil {
		t.Fatal(err)
	}

	// Create two redelegations (pending)
	_, err = s.App.AppKeepers.StakingKeeper.BeginRedelegation(s.Ctx, delegator, srcVal, dstVal0, math.LegacyNewDec(100_000_000))
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.App.AppKeepers.StakingKeeper.BeginRedelegation(s.Ctx, delegator, srcVal, dstVal1, math.LegacyNewDec(150_000_000))
	if err != nil {
		t.Fatal(err)
	}

	// Sanity: redelegations exist
	reds, err := s.App.AppKeepers.StakingKeeper.GetRedelegations(s.Ctx, delegator, 65535)
	if err != nil {
		t.Fatal(err)
	}
	if len(reds) == 0 {
		t.Fatal("expected active redelegations before completion")
	}

	// Act: complete all redelegations at now
	_, err = v5_1.CompleteAllRedelegations(s.Ctx, s.Ctx.BlockTime(), &s.App.AppKeepers, delegator)
	if err != nil {
		t.Fatal(err)
	}

	// Assert: no active redelegations remain
	reds, err = s.App.AppKeepers.StakingKeeper.GetRedelegations(s.Ctx, delegator, 65535)
	if err != nil {
		t.Fatal(err)
	}
	if len(reds) != 0 {
		t.Fatalf("expected 0 redelegations, got %d", len(reds))
	}

	_ = denom // keep for potential future assertions
}

func TestUnbondAllAndFinish(t *testing.T) {
	var s apptesting.TestSuite
	s.Setup()

	delegator := s.TestAccs[0].Address

	// Create one validator and delegate
	val := sdk.ValAddress([]byte("src-validator-addr-__________2"))
	v := ensureValidator(t, &s, val)
	fund(t, &s, delegator, 2_000_000_000)
	_, err := s.App.AppKeepers.StakingKeeper.Delegate(s.Ctx, delegator, math.NewInt(1_000_000_000), stakingtypes.Unbonded, v, true)
	if err != nil {
		t.Fatal(err)
	}

	// Create some unbonding entries via undelegate
	shares := math.LegacyNewDec(100_000_000)
	_, _, err = s.App.AppKeepers.StakingKeeper.Undelegate(s.Ctx, delegator, val, shares)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = s.App.AppKeepers.StakingKeeper.Undelegate(s.Ctx, delegator, val, shares)
	if err != nil {
		t.Fatal(err)
	}

	// Sanity: there should be delegations or UBDs before
	dels, err := s.App.AppKeepers.StakingKeeper.GetAllDelegatorDelegations(s.Ctx, delegator)
	if err != nil {
		t.Fatal(err)
	}
	if len(dels) == 0 {
		t.Fatal("expected delegations before unbond")
	}

	ubd, err := s.App.AppKeepers.StakingKeeper.GetAllUnbondingDelegations(s.Ctx, delegator)
	if err != nil {
		t.Fatal(err)
	}
	if len(ubd) == 0 {
		t.Fatal("expected unbonding delegations entries before finish")
	}

	// Act: unbond all and finish now
	now := s.Ctx.BlockTime().Add(1 * time.Second)
	s.Ctx = s.Ctx.WithBlockTime(now)
	amount, err := v5_1.UnbondAllAndFinish(s.Ctx, now, &s.App.AppKeepers, delegator)
	if err != nil {
		t.Fatal(err)
	}
	if !amount.IsPositive() {
		t.Fatal("expected positive unbonded amount")
	}

	// Assert: no delegations and no unbonding remain for delegator
	dels, err = s.App.AppKeepers.StakingKeeper.GetAllDelegatorDelegations(s.Ctx, delegator)
	if err != nil {
		t.Fatal(err)
	}
	if len(dels) != 0 {
		t.Fatalf("expected 0 delegations, got %d", len(dels))
	}

	ubd, err = s.App.AppKeepers.StakingKeeper.GetAllUnbondingDelegations(s.Ctx, delegator)
	if err != nil {
		t.Fatal(err)
	}
	if len(ubd) != 0 {
		t.Fatalf("expected 0 unbonding delegations, got %d", len(ubd))
	}
}
