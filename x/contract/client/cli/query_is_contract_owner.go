package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/firmachain/firmachain/v05/x/contract/types"
)

var _ = strconv.Itoa(0)

func CmdIsContractOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-contract-owner [fileHash] [ownerAddress]",
		Short: "Check Contract Owner by ownerAddress and fileHash",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			reqFileHash := string(args[0])
			reqOwnerAddress := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.IsContractOwnerRequest{

				FileHash:     string(reqFileHash),
				OwnerAddress: string(reqOwnerAddress),
			}

			res, err := queryClient.IsContractOwner(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
