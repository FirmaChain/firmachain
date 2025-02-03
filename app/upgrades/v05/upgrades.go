package v05

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	//"cosmossdk.io/math"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/firmachain/firmachain/v05/app/keepers"

	//appparamas "github.com/firmachain/firmachain/v05/app/params"

	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v8/types"
	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	// SDK v47 modules
	wasmv2 "github.com/CosmWasm/wasmd/x/wasm/migrations/v2"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"

	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
)

func CreateV0_5_0UpgradeHandler(
	mm *module.Manager,
	cfg module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(c context.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(c)
		logger := ctx.Logger().With("upgrade", UpgradeName)

		// === Params migration ===
		// https://github.com/cosmos/cosmos-sdk/pull/12363/files
		// Set param key table for params module migration
		for _, subspace := range keepers.ParamsKeeper.GetSubspaces() {
			subspace := subspace

			var keyTable paramstypes.KeyTable
			switch subspace.Name() {
			// SDK
			case authtypes.ModuleName:
				keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
			case banktypes.ModuleName:
				keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
			case crisistypes.ModuleName:
				keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck
			case distrtypes.ModuleName:
				keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
			case govtypes.ModuleName:
				keyTable = govv1.ParamKeyTable() //nolint:staticcheck
			case minttypes.ModuleName:
				keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
			case slashingtypes.ModuleName:
				keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
			case stakingtypes.ModuleName:
				keyTable = stakingtypes.ParamKeyTable() //nolint:staticcheck

			// IBC
			case ibctransfertypes.ModuleName:
				keyTable = ibctransfertypes.ParamKeyTable() //nolint:staticcheck
			case icahosttypes.SubModuleName:
				keyTable = icahosttypes.ParamKeyTable() //nolint:staticcheck
			case icacontrollertypes.SubModuleName:
				keyTable = icacontrollertypes.ParamKeyTable() //nolint:staticcheck
			case icqtypes.ModuleName:
				keyTable = icqtypes.ParamKeyTable() //nolint:staticcheck

			// Wasm
			case wasmtypes.ModuleName:
				keyTable = wasmv2.ParamKeyTable() //nolint:staticcheck

			default:
				continue
			}
			if !subspace.HasKeyTable() {
				subspace.WithKeyTable(keyTable)
			}
		}
		// Migrate Tendermint consensus parameters from x/params module to a deprecated x/consensus module.
		// The old params module is required to still be imported in your app.go in order to handle this migration.
		baseAppLegacySS := keepers.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())
		err := baseapp.MigrateParams(ctx, baseAppLegacySS, keepers.ConsensusParamsKeeper.ParamsStore)
		if err != nil {
			return nil, err
		}

		// === New params ===
		// https://github.com/cosmos/ibc-go/blob/v7.1.0/docs/migrations/v7-to-v7_1.md
		// explicitly update the IBC 02-client params, adding the localhost client type
		var newIBCCoreParams clienttypes.Params
		newIBCCoreParams.AllowedClients = append(newIBCCoreParams.AllowedClients,
			ibcexported.Solomachine,
			ibcexported.Tendermint,
			ibcexported.Localhost,
		)
		// ICA Host
		newIcaHostParams := icahosttypes.Params{
			HostEnabled: true,
			// https://github.com/cosmos/ibc-go/blob/v4.2.0/docs/apps/interchain-accounts/parameters.md#allowmessages
			AllowMessages: []string{"*"},
		}
		// ICA Controller
		newIcaControllerParams := icacontrollertypes.Params{ControllerEnabled: true}
		// IBC PFM
		newPFMParams := packetforwardtypes.DefaultParams()
		// ICQ
		newICQParams := icqtypes.NewParams(true, nil)
		// IBC Fee
		logger.Info(fmt.Sprintf("ibcfee module version %s set", fmt.Sprint(vm[ibcfeetypes.ModuleName])))
		// IBC Hooks
		logger.Info(fmt.Sprintf("ibchooks module version %s set", fmt.Sprint(vm[ibchookstypes.ModuleName])))
		// Gov expedited proposal param
		/*
			TODO: check if we want to keep this
			govParams, err := keepers.GovKeeper.Params.Get(ctx)
			if err != nil {
				return nil, err
			}
			govParams.ExpeditedMinDeposit = sdk.NewCoins(sdk.NewCoin(appparamas.BaseCoinUnit, math.NewInt(5000000000)))
			govParams.MinInitialDepositRatio = "0.250000000000000000"
		*/

		// ==== Run migration ====

		logger.Info(fmt.Sprintf("pre migrate version map: %v", vm))
		versionMap, err := mm.RunMigrations(ctx, cfg, vm)
		if err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("post migrate version map: %v", versionMap))

		// New modules run AFTER the migrations, so to set the correct params after the default.

		// ==== Set Params ====
		keepers.IBCKeeper.ClientKeeper.SetParams(ctx, newIBCCoreParams)
		logger.Info(fmt.Sprintf("ibc core: ICQKeeper params set"))

		keepers.ICAHostKeeper.SetParams(ctx, newIcaHostParams)
		logger.Info(fmt.Sprintf("icahost: ICAHostKeeper params set"))

		keepers.ICAControllerKeeper.SetParams(ctx, newIcaControllerParams)
		logger.Info(fmt.Sprintf("icacontroller: ICAControllerKeeper params set"))

		err = keepers.PacketForwardKeeper.SetParams(ctx, newPFMParams)
		if err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("packetforward: PacketForwardKeeper params set"))

		err = keepers.ICQKeeper.SetParams(ctx, newICQParams)
		if err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("icq: ICQKeeper params set"))

		/*
			TODO: uncomment if govParams have to be modified
			err = keepers.GovKeeper.Params.Set(ctx, govParams)
			if err != nil {
				return nil, err
			}
			logger.Info(fmt.Sprintf("icq: GovKeeper params set"))
		*/
		return versionMap, err
	}
}
