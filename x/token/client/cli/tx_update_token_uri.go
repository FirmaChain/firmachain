package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/firmachain/firmachain/x/token/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateTokenURI() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-token-uri [token-id] [token-uri]",
		Short: "Broadcast message updateTokenURI",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenID := args[0]
			argTokenURI := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTokenURI(
				clientCtx.GetFromAddress().String(),
				argTokenID,
				argTokenURI,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
