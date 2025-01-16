package main

import (
	"os"

	"cosmossdk.io/log"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/firmachain/firmachain/v05/app"
)

func main() {

	rootCmd := NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "FIRMACHAIND", app.DefaultNodeHome); err != nil {
		log.NewLogger(rootCmd.OutOrStderr()).Error("failure when running app", "err", err)
		os.Exit(1)

	}
}
