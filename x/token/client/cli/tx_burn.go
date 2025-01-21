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

func CmdBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [token-id] [amount]",
		Short: "Broadcast message burn",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenID := args[0]
			argAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argAmountValue, err := strconv.ParseUint(argAmount, 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurn(
				clientCtx.GetFromAddress().String(),
				argTokenID,
				argAmountValue,
			)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
