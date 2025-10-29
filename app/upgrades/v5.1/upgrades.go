package v5_1

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/math"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/firmachain/firmachain/app/keepers"

	"github.com/cosmos/cosmos-sdk/codec"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	appparams "github.com/firmachain/firmachain/app/params"
)

var oldValidators = []string{
	"firmavaloper1p90hu6pqd57xgdauf0l8dwpau73jkk7273y6ck",
	"firmavaloper1x0lqg5vcynse3r6mug8vteu77cqyaqkgw2rg3x",
	"firmavaloper1cztkfn8alkj69d79hvygvdueuvyxem7zvu2s9e",
	"firmavaloper1vqr622jkes6xmh2cxqa7a04dp6seyme3gkxml4",
	"firmavaloper1pa8yzyytgfg5hk75gmhkpcq6kds57ak39ydw8v",
	"firmavaloper1d84pmnumnsh80v74lta0vnpd476ncp4pvqhdld",
	"firmavaloper1u7f4kz740jcr5skq9xnpw8568uugtja5u0k87v",
	"firmavaloper1paujadj4fxmqrzwrnhqacuk5dst3fkhc5gwddu",
	"firmavaloper16treuwa6d64xs4r4sshurw75363pwp4n86s6sq",
	"firmavaloper1q90jttne2mje828mfh3nmfw8kucdsfwf6a5c20",
	"firmavaloper1tjc2j8aq59s2t9waakacj9xc76yx97t7c55hkh",
}

type accountMove struct {
	oldAccStr string
	newAccStr string
	newValStr string
}

var moves = []accountMove{
	{
		oldAccStr: "firma1p90hu6pqd57xgdauf0l8dwpau73jkk72qz0pcc",
		newAccStr: "firma1k0m54qycp4v04wazj0f72snp86htau5g6ujfau",
		newValStr: "firmavaloper1k0m54qycp4v04wazj0f72snp86htau5gy0ejaj",
		// TODO: toMoveBal: 12345...
	},
	{
		oldAccStr: "firma1x0lqg5vcynse3r6mug8vteu77cqyaqkgsegn3g",
		newAccStr: "firma1fzupf3r5gk505ddt4qpms0gsa7e09j68kzexzl",
		newValStr: "firmavaloper1fzupf3r5gk505ddt4qpms0gsa7e09j68g3jaz3",
	},
	{
		oldAccStr: "firma1cztkfn8alkj69d79hvygvdueuvyxem7zj0pt9h",
		newAccStr: "firma1c2d74r4fgav52tuq9w9e3sz6d7dk2hx6sgvarp",
		newValStr: "firmavaloper1c2d74r4fgav52tuq9w9e3sz6d7dk2hx6wm8xr0",
	},
	{
		oldAccStr: "firma1vqr622jkes6xmh2cxqa7a04dp6seyme3k9dqlm",
		newAccStr: "firma1c2d74r4fgav52tuq9w9e3sz6d7dk2hx6sgvarp",
		newValStr: "firmavaloper1c2d74r4fgav52tuq9w9e3sz6d7dk2hx6wm8xr0",
	},
	{
		oldAccStr: "firma1pa8yzyytgfg5hk75gmhkpcq6kds57ak3mhx48z",
		newAccStr: "firma1wer4d5h7rs0knhcf7jcgevjqymmx3szyf8ptgc",
		newValStr: "firmavaloper1wer4d5h7rs0knhcf7jcgevjqymmx3szyh52sgk",
	},
	{
		oldAccStr: "firma1d84pmnumnsh80v74lta0vnpd476ncp4pjnuklr",
		newAccStr: "firma1wer4d5h7rs0knhcf7jcgevjqymmx3szyf8ptgc",
		newValStr: "firmavaloper1wer4d5h7rs0knhcf7jcgevjqymmx3szyh52sgk",
	},
	{
		oldAccStr: "firma1u7f4kz740jcr5skq9xnpw8568uugtja5zuau7z",
		newAccStr: "firma1q3vad3h69pe8xe4lk3uqdpy87tj3hydyqhtfzw",
		newValStr: "firmavaloper1q3vad3h69pe8xe4lk3uqdpy87tj3hydy7yqjzq",
	},
	{
		oldAccStr: "firma1paujadj4fxmqrzwrnhqacuk5dst3fkhc2m9kdj",
		newAccStr: "firma1q3vad3h69pe8xe4lk3uqdpy87tj3hydyqhtfzw",
		newValStr: "firmavaloper1q3vad3h69pe8xe4lk3uqdpy87tj3hydy7yqjzq",
	},
	{
		oldAccStr: "firma16treuwa6d64xs4r4sshurw75363pwp4nefmpsw",
		newAccStr: "firma1qy5rfrzn9jtclncyx3c9fp5ea85ak0uv4pvudj",
		newValStr: "firmavaloper1qy5rfrzn9jtclncyx3c9fp5ea85ak0uvtj88du",
	},
	{
		oldAccStr: "firma1q90jttne2mje828mfh3nmfw8kucdsfwfywlr2p",
		newAccStr: "firma1qy5rfrzn9jtclncyx3c9fp5ea85ak0uv4pvudj",
		newValStr: "firmavaloper1qy5rfrzn9jtclncyx3c9fp5ea85ak0uvtj88du",
	},
	{
		oldAccStr: "firma1tjc2j8aq59s2t9waakacj9xc76yx97t7x8lvke",
		newAccStr: "firma1qy5rfrzn9jtclncyx3c9fp5ea85ak0uv4pvudj",
		newValStr: "firmavaloper1qy5rfrzn9jtclncyx3c9fp5ea85ak0uvtj88du",
	},
}

func CreateV0_5_1UpgradeHandler(
	mm *module.Manager,
	cfg module.Configurator,
	keepers *keepers.AppKeepers,
	appCodec *codec.ProtoCodec,
) upgradetypes.UpgradeHandler {
	return func(c context.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(c)
		logger := ctx.Logger().With("upgrade", UpgradeName)

		// ==== Run migration ====
		logger.Info(fmt.Sprintf("pre migrate version map: %v", vm))
		versionMap, err := mm.RunMigrations(ctx, cfg, vm)
		if err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("post migrate version map: %v", versionMap))

		// ==== Custom migration: account moves and delegations ====
		bondDenom := appparams.DefaultBondDenom

		// Withdraw commission of all old validators
		for _, valAddrStr := range oldValidators {
			valAddr, err := sdk.ValAddressFromBech32(valAddrStr)
			if err != nil {
				return nil, fmt.Errorf("invalid validator address %s: %w", valAddr, err)
			}

			if _, err := keepers.DistrKeeper.WithdrawValidatorCommission(ctx, valAddr); err != nil { // TODO: if we move only toMoveAmount, these commissions are not moved. How to handle?
				logger.Error("withdraw validator commission failed", "validator", valAddr, "err", err)
			}
		}

		// Move all accounts
		for _, mv := range moves {
			oldAcc, err := sdk.AccAddressFromBech32(mv.oldAccStr)
			if err != nil {
				return nil, fmt.Errorf("invalid old acc address %s: %w", mv.oldAccStr, err)
			}
			newAcc, err := sdk.AccAddressFromBech32(mv.newAccStr)
			if err != nil {
				return nil, fmt.Errorf("invalid new acc address %s: %w", mv.newAccStr, err)
			}
			newValAddr, err := sdk.ValAddressFromBech32(mv.newValStr)
			if err != nil {
				return nil, fmt.Errorf("invalid new validator address %s: %w", mv.newValStr, err)
			}

			// Complete all redelegations
			redelegatedAmt, err := CompleteAllRedelegations(ctx, ctx.BlockTime(), keepers, oldAcc)
			if err != nil {
				return nil, fmt.Errorf("complete all redelegations failed: %w", err)
			}
			logger.Info("redelegated amount", "amount", redelegatedAmt.String())

			// Unbond all and terminate
			unbondedAmt, err := UnbondAllAndFinish(ctx, ctx.BlockTime(), keepers, oldAcc)
			if err != nil {
				return nil, fmt.Errorf("unbond all and finish failed: %w", err)
			}
			logger.Info("unbonded amount", "amount", unbondedAmt.String())

			// Move all bank balances from old account to new account
			transferCoins := keepers.BankKeeper.GetAllBalances(ctx, oldAcc)
			if !transferCoins.IsZero() {
				// TODO: why tErr?
				// TODO: dont transfer transferCoins but toMoveAmount (slightly less than the whole)
				if tErr := keepers.BankKeeper.SendCoins(ctx, oldAcc, newAcc, transferCoins); tErr != nil {
					return nil, tErr
				}
				logger.Info("migrated balances", "from", mv.oldAccStr, "to", mv.newAccStr, "coins", transferCoins.String())
			}

			// Delegate the transferred bond-denom amount from the new account to the new validator, keep 10 FCT as reserve
			// TODO: don't delegate
			bondAmt := transferCoins.AmountOf(bondDenom).Sub(math.NewInt(10000000))
			if bondAmt.IsPositive() {
				validator, gErr := keepers.StakingKeeper.GetValidator(ctx, newValAddr) // TODO: why gErr?
				if gErr != nil {
					return nil, gErr
				}
				if _, dErr := keepers.StakingKeeper.Delegate(ctx, newAcc, bondAmt, stakingtypes.Unbonded, validator, true); dErr != nil { // TODO: why dErr?
					return nil, dErr
				}
				logger.Info("delegated migrated stake", "delegator", mv.newAccStr, "validator", mv.newValStr, "amount", bondAmt.String(), "denom", bondDenom)
			}
		}

		return versionMap, err
	}
}
