package v5_1

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/firmachain/firmachain/app/keepers"
)

func CreateV0_5_1UpgradeHandler(
	mm *module.Manager,
	cfg module.Configurator,
	_ *keepers.AppKeepers,
	_ *codec.ProtoCodec,
) upgradetypes.UpgradeHandler {
	return func(c context.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(c)
		return mm.RunMigrations(ctx, cfg, vm)
	}
}
