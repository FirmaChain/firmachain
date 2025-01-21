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

func CmdMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [tokenUri]",
		Short: "Broadcast message mint",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsTokenURI := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(clientCtx.GetFromAddress().String(), string(argsTokenURI))
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
