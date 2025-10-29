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
		return redelegatedAmt, err
	}
	for _, activeRedelegation := range redelegations {
		redelegationSrc, _ := sdk.ValAddressFromBech32(activeRedelegation.ValidatorSrcAddress)
		redelegationDst, _ := sdk.ValAddressFromBech32(activeRedelegation.ValidatorDstAddress)

		// set all entry completionTime to now so we can complete re-delegation
		for i := range activeRedelegation.Entries {
			activeRedelegation.Entries[i].CompletionTime = now
			redelegatedAmt = redelegatedAmt.Add(math.Int(activeRedelegation.Entries[i].SharesDst)) // TODO: Shares =/= Amount
		}

		keepers.StakingKeeper.SetRedelegation(ctx, activeRedelegation)
		_, err = keepers.StakingKeeper.CompleteRedelegation(ctx, accAddr, redelegationSrc, redelegationDst) // TODO: Use this for the redelegatedAmt
		if err != nil {
			return redelegatedAmt, err
		}
	}

	return redelegatedAmt, nil
}

// Returns the amount of tokens which were unbonded (not rewards)
func UnbondAllAndFinish(ctx sdk.Context, now time.Time, keepers *keepers.AppKeepers, accAddr sdk.AccAddress) (math.Int, error) {
	unbondedAmt := math.ZeroInt()

	// Unbond all delegations from the account
	delegations, err := keepers.StakingKeeper.GetAllDelegatorDelegations(ctx, accAddr)
	if err != nil {
		return math.ZeroInt(), err
	}
	for _, delegation := range delegations {

		validatorValAddr, err := sdk.ValAddressFromBech32(delegation.GetValidatorAddr())
		if err != nil {
			continue // TODO: make it fail
		}

		if _, err := keepers.StakingKeeper.GetValidator(ctx, validatorValAddr); err != nil {
			continue // TODO: make it fail
		}

		// TODO: 1. make clearSelfUnbond function
		// TODO: 2. use clearSelfUnbond here to clear the current self-unbonding queue
		// TODO: 3. leave a sufficient self-delegation: use undelegateAmount:=delegation.GetShares()-MIN_SELF_DELEGATION ?
		_, _, err = keepers.StakingKeeper.Undelegate(ctx, accAddr, validatorValAddr, delegation.GetShares()) // TODO: what if validator has already 7 entries in its unbonding-delegations queue?
		if err != nil {
			return math.ZeroInt(), err
		}
	}

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
