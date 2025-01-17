package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateTokenUri() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-token-uri [token-id] [token-uri]",
		Short: "Broadcast message updateTokenUri",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenId := args[0]
			argTokenUri := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTokenUri(
				clientCtx.GetFromAddress().String(),
				argTokenId,
				argTokenUri,
			)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
