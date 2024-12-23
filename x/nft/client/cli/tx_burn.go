package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/firmachain/firmachain/v05/x/nft/types"
)

var _ = strconv.Itoa(0)

func CmdBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [nftId]",
		Short: "Broadcast message burn",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsNftId, err := strconv.ParseUint(args[0], 10, 64)

			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurn(clientCtx.GetFromAddress().String(), argsNftId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
