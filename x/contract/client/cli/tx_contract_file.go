package cli

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/firmachain/firmachain/v05/x/contract/types"
)

func CmdCreateContractFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-contract-file [fileHash] [timeStamp] [ownerList(\"address,address\")] [metaDataJsonString]",
		Short: "Create a new contractFile",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			fileHash := args[0]
			argsTimeStamp, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argsOwnerJsonString, err := cast.ToStringE(args[2])

			argsOwnerList := strings.Split(argsOwnerJsonString, ",")

			for i := 0; i < len(argsOwnerList); i++ {
				argsOwnerList[i] = strings.TrimSpace(argsOwnerList[i])
			}

			if err != nil {
				return err
			}
			argsMetaDataJsonString, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateContractFile(clientCtx.GetFromAddress().String(), fileHash, argsTimeStamp, argsOwnerList, argsMetaDataJsonString)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
