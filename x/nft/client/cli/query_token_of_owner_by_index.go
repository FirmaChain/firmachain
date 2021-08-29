package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/firmachain/firmachain/x/nft/types"
)

var _ = strconv.Itoa(0)

func CmdTokenOfOwnerByIndex() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-of-owner-by-index [ownerAddress] [index]",
		Short: "Query tokenOfOwnerByIndex",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			reqOwnerAddress := string(args[0])
			reqIndex, _ := strconv.ParseUint(args[1], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryTokenOfOwnerByIndexRequest{

				OwnerAddress: string(reqOwnerAddress),
				Index:        uint64(reqIndex),
			}

			res, err := queryClient.TokenOfOwnerByIndex(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
