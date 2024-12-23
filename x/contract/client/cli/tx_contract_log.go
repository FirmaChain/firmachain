package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/firmachain/firmachain/v05/x/contract/types"
)

func CmdAddContractLog() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-contract-log [contractHash] [timeStamp] [eventName] [ownerAddress] [jsonString]",
		Short: "Add a new contractLog",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsContractHash, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsTimeStamp, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argsEventName, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsOwnerAddress, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			argsJsonString, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddContractLog(clientCtx.GetFromAddress().String(), argsContractHash, argsTimeStamp, argsEventName, argsOwnerAddress, argsJsonString)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
