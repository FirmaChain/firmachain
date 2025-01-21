package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/firmachain/firmachain/v05/x/contract/types"
)

var _ = strconv.Itoa(0)

func CmdContractListFromHash() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-contract-list-from-hash [contractHash]",
		Short: "Get Contract from contract hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			reqContractHash := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetContractListFromHashRequest{
				Hash: string(reqContractHash),
			}

			res, err := queryClient.GetContractListFromHash(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
