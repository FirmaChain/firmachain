package cli

import (
	"bufio"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/firmachain/FirmaChain/x/contract/types"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "contract transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(flags.PostCommands(GetCmdAddContract(cdc))...)

	return cmd
}

func GetCmdAddContract(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add [path] [hash]",
		Short: "Add contract to blockchain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

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
