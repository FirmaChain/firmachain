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
	"firmavaloper1cztkfn8alkj69d79hvygvdueuvyxem7zvu2s9e",
	"firmavaloper1vqr622jkes6xmh2cxqa7a04dp6seyme3gkxml4",
}

type AccountMove struct {
	OldAccStr string
	NewAccStr string
	NewValStr string
}

var Moves = []AccountMove{
	{
		OldAccStr: "firma1cztkfn8alkj69d79hvygvdueuvyxem7zj0pt9h",
		NewAccStr: "firma1c2d74r4fgav52tuq9w9e3sz6d7dk2hx6sgvarp",
		NewValStr: "firmavaloper1nssuz67am2uwc2hjgvphg0fmj3k9l6cxy8ha9j",
	},
	{
		OldAccStr: "firma1vqr622jkes6xmh2cxqa7a04dp6seyme3k9dqlm",
		NewAccStr: "firma1c2d74r4fgav52tuq9w9e3sz6d7dk2hx6sgvarp",
		NewValStr: "firmavaloper1nssuz67am2uwc2hjgvphg0fmj3k9l6cxy8ha9j",
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
		if ctx.ChainID() == "imperium-4" { // testnet-only version
			bondDenom := appparams.DefaultBondDenom

			// Withdraw commission of all old validators
			for _, valAddrStr := range OldValidators {
				valAddr, err := sdk.ValAddressFromBech32(valAddrStr)
				if err != nil {
					return nil, fmt.Errorf("invalid validator address %s: %w", valAddr, err)
				}

				err = keepers.DistrKeeper.SetWithdrawAddr(ctx, sdk.AccAddress(valAddr), sdk.AccAddress(valAddr))
				if err != nil {
					return nil, err
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
