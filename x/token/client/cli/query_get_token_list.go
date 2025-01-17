package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

var _ = strconv.Itoa(0)

func CmdGetTokenList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-token-list [owner-address]",
		Short: "Query getTokenList",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOwnerAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.GetTokenListRequest{

				OwnerAddress: reqOwnerAddress,
			}

			res, err := queryClient.GetTokenList(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
