package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/firmachain/firmachain/v05/x/contract/types"
	"github.com/spf13/cobra"
)

func CmdListContractFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-contract-file",
		Short: "list all contractFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllContractFileRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ContractFileAll(context.Background(), params)
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

func CmdShowContractFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-contract-file [index]",
		Short: "shows a contractFile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetContractFileRequest{
				Index: args[0],
			}

			res, err := queryClient.ContractFile(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
