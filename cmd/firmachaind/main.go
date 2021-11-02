package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/firmachain/firmachain/app"
	"github.com/firmachain/firmachain/spm/cosmoscmd"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
	)

	// below 2 codes use for command line description by starport and makefile.
	rootCmd.Use = "firmachaind"
	rootCmd.Short = "\n FirmaChain BlockChain [https://firmachain.org]"

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
