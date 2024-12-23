package cli

import (
	"strconv"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

var _ = strconv.Itoa(0)

func CmdCreateToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-token [name] [symbol] [token-uri] [total-supply] [decimal] [mintable] [burnable]",
		Short: "Broadcast message createToken",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argSymbol := args[1]
			argTokenURI := args[2]
			argTotalSupply := args[3]
			argDecimal := args[4]
			argMintable, err := cast.ToBoolE(args[5])
			if err != nil {
				return err
			}
			argBurnable, err := cast.ToBoolE(args[6])
			if err != nil {
				return err
			}

			argTotalSupplyValue, err := strconv.ParseUint(argTotalSupply, 10, 64)
			if err != nil {
				return err
			}

			argDecimalValue, err := strconv.ParseUint(argDecimal, 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateToken(
				clientCtx.GetFromAddress().String(),
				argName,
				argSymbol,
				argTokenURI,
				argTotalSupplyValue,
				argDecimalValue,
				argMintable,
				argBurnable,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
