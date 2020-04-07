package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/firmachain/FirmaChain/x/contract/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	FirmaChainTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "FirmaChain transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	FirmaChainTxCmd.AddCommand(client.PostCommands(
		GetCmdAddContract(cdc),
	)...)

	return FirmaChainTxCmd
}
func GetCmdAddContract(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-contract [path] [hash]",
		Short: "add-contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgAddContract(args[0], args[1], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			broadcast := utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})

			return broadcast
		},
	}
}
