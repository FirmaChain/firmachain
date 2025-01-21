package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/firmachain/firmachain/v05/x/nft/types"
)

var _ = strconv.Itoa(0)

func CmdNftIdListOfOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-nft-id-of-owner [ownerAddress]",
		Short: "Query Nft ID List Of Owner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			reqOwnerAddress := string(args[0])

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryNftIdListOfOwnerRequest{

				OwnerAddress: string(reqOwnerAddress),
				Pagination:   pageReq,
			}

			res, err := queryClient.NftIdListOfOwner(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
