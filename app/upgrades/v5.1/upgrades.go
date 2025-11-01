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

var OldValidators = []string{
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

type AccountMove struct {
	OldAccStr string
	NewAccStr string
	NewValStr string
}

var Moves = []AccountMove{
	{
		OldAccStr: "firma1p90hu6pqd57xgdauf0l8dwpau73jkk72qz0pcc",
		NewAccStr: "firma1k0m54qycp4v04wazj0f72snp86htau5g6ujfau",
		NewValStr: "firmavaloper1k0m54qycp4v04wazj0f72snp86htau5gy0ejaj",
	},
	{
		OldAccStr: "firma1x0lqg5vcynse3r6mug8vteu77cqyaqkgsegn3g",
		NewAccStr: "firma1fzupf3r5gk505ddt4qpms0gsa7e09j68kzexzl",
		NewValStr: "firmavaloper1fzupf3r5gk505ddt4qpms0gsa7e09j68g3jaz3",
	},
	{
		OldAccStr: "firma1cztkfn8alkj69d79hvygvdueuvyxem7zj0pt9h",
		NewAccStr: "firma1c2d74r4fgav52tuq9w9e3sz6d7dk2hx6sgvarp",
		NewValStr: "firmavaloper1c2d74r4fgav52tuq9w9e3sz6d7dk2hx6wm8xr0",
	},
	{
		OldAccStr: "firma1vqr622jkes6xmh2cxqa7a04dp6seyme3k9dqlm",
		NewAccStr: "firma1c2d74r4fgav52tuq9w9e3sz6d7dk2hx6sgvarp",
		NewValStr: "firmavaloper1c2d74r4fgav52tuq9w9e3sz6d7dk2hx6wm8xr0",
	},
	{
		OldAccStr: "firma1pa8yzyytgfg5hk75gmhkpcq6kds57ak3mhx48z",
		NewAccStr: "firma1wer4d5h7rs0knhcf7jcgevjqymmx3szyf8ptgc",
		NewValStr: "firmavaloper1wer4d5h7rs0knhcf7jcgevjqymmx3szyh52sgk",
	},
	{
		OldAccStr: "firma1d84pmnumnsh80v74lta0vnpd476ncp4pjnuklr",
		NewAccStr: "firma1wer4d5h7rs0knhcf7jcgevjqymmx3szyf8ptgc",
		NewValStr: "firmavaloper1wer4d5h7rs0knhcf7jcgevjqymmx3szyh52sgk",
	},
	{
		OldAccStr: "firma1u7f4kz740jcr5skq9xnpw8568uugtja5zuau7z",
		NewAccStr: "firma1q3vad3h69pe8xe4lk3uqdpy87tj3hydyqhtfzw",
		NewValStr: "firmavaloper1q3vad3h69pe8xe4lk3uqdpy87tj3hydy7yqjzq",
	},
	{
		OldAccStr: "firma1paujadj4fxmqrzwrnhqacuk5dst3fkhc2m9kdj",
		NewAccStr: "firma1q3vad3h69pe8xe4lk3uqdpy87tj3hydyqhtfzw",
		NewValStr: "firmavaloper1q3vad3h69pe8xe4lk3uqdpy87tj3hydy7yqjzq",
	},
	{
		OldAccStr: "firma16treuwa6d64xs4r4sshurw75363pwp4nefmpsw",
		NewAccStr: "firma1qy5rfrzn9jtclncyx3c9fp5ea85ak0uv4pvudj",
		NewValStr: "firmavaloper1qy5rfrzn9jtclncyx3c9fp5ea85ak0uvtj88du",
	},
	{
		OldAccStr: "firma1q90jttne2mje828mfh3nmfw8kucdsfwfywlr2p",
		NewAccStr: "firma1qy5rfrzn9jtclncyx3c9fp5ea85ak0uv4pvudj",
		NewValStr: "firmavaloper1qy5rfrzn9jtclncyx3c9fp5ea85ak0uvtj88du",
	},
	{
		OldAccStr: "firma1tjc2j8aq59s2t9waakacj9xc76yx97t7x8lvke",
		NewAccStr: "firma1qy5rfrzn9jtclncyx3c9fp5ea85ak0uv4pvudj",
		NewValStr: "firmavaloper1qy5rfrzn9jtclncyx3c9fp5ea85ak0uvtj88du",
	},
}

var NotDelegatedAmount = int64(10_000_000) // ufct

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

		// ==== Custom migration: account Moves and delegations ====
		if ctx.ChainID() == "colosseum-1" {
			bondDenom := appparams.DefaultBondDenom

			// Withdraw commission of all old validators
			for _, valAddrStr := range OldValidators {
				valAddr, err := sdk.ValAddressFromBech32(valAddrStr)
				if err != nil {
					return nil, fmt.Errorf("invalid validator address %s: %w", valAddr, err)
				}

				withdrawAddr, err := keepers.DistrKeeper.GetDelegatorWithdrawAddr(ctx, sdk.AccAddress(valAddr))
				if err != nil {
					err = keepers.DistrKeeper.SetWithdrawAddr(ctx, sdk.AccAddress(valAddr), sdk.AccAddress(valAddr))
					if err != nil {
						return nil, err
					}
				} else {
					err = keepers.DistrKeeper.DeleteDelegatorWithdrawAddr(ctx, sdk.AccAddress(valAddr), withdrawAddr)
					if err != nil {
						return nil, err
					}
					err = keepers.DistrKeeper.SetWithdrawAddr(ctx, sdk.AccAddress(valAddr), sdk.AccAddress(valAddr))
					if err != nil {
						return nil, err
					}
				}

				if _, err := keepers.DistrKeeper.WithdrawValidatorCommission(ctx, valAddr); err != nil {
					logger.Error("withdraw validator commission failed", "validator", valAddr, "err", err)
				}
			}

			// Move all accounts
			for _, mv := range Moves {
				oldAcc, err := sdk.AccAddressFromBech32(mv.OldAccStr)
				if err != nil {
					return nil, fmt.Errorf("invalid old acc address %s: %w", mv.OldAccStr, err)
				}
				newAcc, err := sdk.AccAddressFromBech32(mv.NewAccStr)
				if err != nil {
					return nil, fmt.Errorf("invalid new acc address %s: %w", mv.NewAccStr, err)
				}
				newValAddr, err := sdk.ValAddressFromBech32(mv.NewValStr)
				if err != nil {
					return nil, fmt.Errorf("invalid new validator address %s: %w", mv.NewValStr, err)
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
					if tErr := keepers.BankKeeper.SendCoins(ctx, oldAcc, newAcc, transferCoins); tErr != nil {
						return nil, tErr
					}
					logger.Info("migrated balances", "from", mv.OldAccStr, "to", mv.NewAccStr, "coins", transferCoins.String())
				}

				// Delegate all the current balance (transfered amount plus initial amount), keeping 10 FCT as reserve
				bondAmt := keepers.BankKeeper.GetBalance(ctx, newAcc, bondDenom).Amount.Sub(math.NewInt(NotDelegatedAmount))
				if bondAmt.IsPositive() {
					validator, gErr := keepers.StakingKeeper.GetValidator(ctx, newValAddr)
					if gErr != nil {
						return nil, gErr
					}
					if _, dErr := keepers.StakingKeeper.Delegate(ctx, newAcc, bondAmt, stakingtypes.Unbonded, validator, true); dErr != nil {
						return nil, dErr
					}
					logger.Info("delegated migrated stake", "delegator", mv.NewAccStr, "validator", mv.NewValStr, "amount", bondAmt.String(), "denom", bondDenom)
				}
			}
		}

		return versionMap, err
	}
}
