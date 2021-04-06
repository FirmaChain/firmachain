package main

import (
	"encoding/json"
	app2 "github.com/firmachain/FirmaChain/app"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	"github.com/firmachain/FirmaChain/types/address"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

// firma custom flags
const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cobra.EnableCommandSorting = false

	cdc := app2.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(address.Bech32PrefixAccAddr, address.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(address.Bech32PrefixValAddr, address.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(address.Bech32PrefixConsAddr, address.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "firma",
		Short:             "FirmaChain (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(
		genutilcli.InitCmd(ctx, cdc, app2.ModuleBasics, app2.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(ctx, cdc, auth.GenesisAccountIterator{}, app2.DefaultNodeHome),
		genutilcli.GenTxCmd(
			ctx, cdc, app2.ModuleBasics, staking.AppModuleBasic{},
			auth.GenesisAccountIterator{}, app2.DefaultNodeHome, app2.DefaultCLIHome,
		),
		genutilcli.ValidateGenesisCmd(ctx, cdc, app2.ModuleBasics),
		AddGenesisAccountCmd(ctx, cdc, app2.DefaultNodeHome, app2.DefaultCLIHome),
	)

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	executor := cli.PrepareBaseCmd(rootCmd, "FC", app2.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range viper.GetIntSlice(server.FlagUnsafeSkipUpgrades) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags()
	if err != nil {
		panic(err)
	}

	return app2.NewFirmaChainApp(
		logger, db, traceStore,
		true, invCheckPeriod, skipUpgradeHeights,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
		baseapp.SetHaltTime(uint64(viper.GetInt(server.FlagHaltTime))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		fcApp := app2.NewFirmaChainApp(logger, db, traceStore, false, uint(1), map[int64]bool{})
		err := fcApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return fcApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	fcApp := app2.NewFirmaChainApp(logger, db, traceStore, true, uint(1), map[int64]bool{})

	return fcApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
