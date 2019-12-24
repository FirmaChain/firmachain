package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/firmachain/FirmaChain/x/firmachain/internal/types"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	FirmaChainQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the firmachain module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	FirmaChainQueryCmd.AddCommand(client.GetCommands(
		GetContract(storeKey, cdc),
	)...)
	return FirmaChainQueryCmd
}

func GetContract(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "contract [hash]",
		Short: "Query contract info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			hash := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/contract/%s", queryRoute, hash), nil)
			if err != nil {
				fmt.Printf("could not resolve hash - %s \n", hash)
				return nil
			}

			var out types.Contract
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
