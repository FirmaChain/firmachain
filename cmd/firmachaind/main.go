package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/firmachain/firmachain/app"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
)

func main() {
	rootCmd, encodingConfig := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
	)

	rootCmd.AddCommand(server.RosettaCommand(encodingConfig.InterfaceRegistry, encodingConfig.Marshaler), AddGenesisWasmMsgCmd(app.DefaultNodeHome))

	// below 2 codes use for command line description by starport and makefile.
	rootCmd.Use = "firmachaind"
	rootCmd.Short = "\n FirmaChain BlockChain [https://firmachain.org]"

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
