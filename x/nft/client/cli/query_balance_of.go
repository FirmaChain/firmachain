package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/firmachain/firmachain/v05/x/nft/types"
)

var _ = strconv.Itoa(0)

func CmdBalanceOf() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance-of [ownerAddress]",
		Short: "Query balanceOf",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			reqOwnerAddress := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryBalanceOfRequest{

				OwnerAddress: string(reqOwnerAddress),
			}

			res, err := queryClient.BalanceOf(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
