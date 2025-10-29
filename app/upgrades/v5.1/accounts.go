package v5_1

import (
	"time"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/app/keepers"
)

// Completes all re-delegations and returns the amount of tokens which were re-delegated.
func CompleteAllRedelegations(ctx sdk.Context, now time.Time, keepers *keepers.AppKeepers, accAddr sdk.AccAddress) (math.Int, error) {
	redelegatedAmt := math.ZeroInt()

	redelegations, err := keepers.StakingKeeper.GetRedelegations(ctx, accAddr, 65535)
	if err != nil {
		return math.ZeroInt(), err
	}
	for _, activeRedelegation := range redelegations {
		redelegationSrc, _ := sdk.ValAddressFromBech32(activeRedelegation.ValidatorSrcAddress)
		redelegationDst, _ := sdk.ValAddressFromBech32(activeRedelegation.ValidatorDstAddress)

		// set all entry completionTime to now so we can complete re-delegation
		for i := range activeRedelegation.Entries {
			activeRedelegation.Entries[i].CompletionTime = now
		}

		keepers.StakingKeeper.SetRedelegation(ctx, activeRedelegation)
		redelegatedCoins, err := keepers.StakingKeeper.CompleteRedelegation(ctx, accAddr, redelegationSrc, redelegationDst)
		if err != nil {
			return math.ZeroInt(), err
		}
		for i := range redelegatedCoins {
			redelegatedAmt = redelegatedAmt.Add(redelegatedCoins[i].Amount)
		}

	}

	return redelegatedAmt, nil
}

// Returns the amount of tokens which were unbonded (not rewards)
func UnbondAllAndFinish(ctx sdk.Context, now time.Time, keepers *keepers.AppKeepers, accAddr sdk.AccAddress) (math.Int, error) {
	unbondedAmt := math.ZeroInt()

	// First clear current unbonding delegations, in order to ensure to have space for the new ones.
	unbondedAmt0, err := CompleteAllUnbondingDelegations(ctx, now, keepers, accAddr)
	if err != nil {
		return math.ZeroInt(), err
	}
	unbondedAmt = unbondedAmt.Add(unbondedAmt0)

	// Unbond all delegations from the account
	delegations, err := keepers.StakingKeeper.GetAllDelegatorDelegations(ctx, accAddr)
	if err != nil {
		return math.ZeroInt(), err
	}
	for _, delegation := range delegations {

		validatorValAddr, err := sdk.ValAddressFromBech32(delegation.GetValidatorAddr())
		if err != nil {
			return math.ZeroInt(), err
		}

		if _, err := keepers.StakingKeeper.GetValidator(ctx, validatorValAddr); err != nil {
			return math.ZeroInt(), err
		}

		_, _, err = keepers.StakingKeeper.Undelegate(ctx, accAddr, validatorValAddr, delegation.GetShares())
		if err != nil {
			return math.ZeroInt(), err
		}
	}

	// Take all unbonding and complete them.
	unbondedAmt1, err := CompleteAllUnbondingDelegations(ctx, now, keepers, accAddr)
	if err != nil {
		return math.ZeroInt(), err
	}
	unbondedAmt = unbondedAmt.Add(unbondedAmt1)

	return unbondedAmt, nil
}

// Completes all unbonding delegations and returns the (final) amount of tokens which are returned from the unbond.
func CompleteAllUnbondingDelegations(ctx sdk.Context, now time.Time, keepers *keepers.AppKeepers, accAddr sdk.AccAddress) (math.Int, error) {
	unbondedAmt := math.ZeroInt()

	// Take all unbonding and complete them.
	unbondingDelegations, err := keepers.StakingKeeper.GetAllUnbondingDelegations(ctx, accAddr)
	if err != nil {
		return math.ZeroInt(), err
	}
	for _, unbondingDelegation := range unbondingDelegations {
		validatorStringAddr := unbondingDelegation.ValidatorAddress
		validatorValAddr, _ := sdk.ValAddressFromBech32(validatorStringAddr)

		// Complete unbonding delegation
		for i := range unbondingDelegation.Entries {
			unbondingDelegation.Entries[i].CompletionTime = now
			unbondedAmt = unbondedAmt.Add(unbondingDelegation.Entries[i].Balance)
		}

		keepers.StakingKeeper.SetUnbondingDelegation(ctx, unbondingDelegation)
		_, err := keepers.StakingKeeper.CompleteUnbonding(ctx, accAddr, validatorValAddr)
		if err != nil {
			return math.ZeroInt(), err
		}
	}

	return unbondedAmt, nil
}
