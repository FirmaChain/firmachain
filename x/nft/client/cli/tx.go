package cli

import (
	"bufio"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/firmachain/FirmaChain/x/nft/types"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "NFT transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(flags.PostCommands(GetCmdAddNFToken(cdc))...)
	cmd.AddCommand(flags.PostCommands(GetCmdTransferNFToken(cdc))...)

	return cmd
}

func GetCmdAddNFToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add [hash] [tokenURI]",
		Short: "Add new NFT to blockchain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			msg := types.NewMsgAddNFToken(args[0], args[1], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()

			if err != nil {
				return err
			}

			broadcast := utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})

			return broadcast
		},
	}
}

func GetCmdTransferNFToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "transfer [hash] [recipient]",
		Short: "Transfer NFT to recipient",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferNFToken(args[0], cliCtx.GetFromAddress(), addr)
			err = msg.ValidateBasic()

			if err != nil {
				return err
			}

			broadcast := utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})

			return broadcast
		},
	}
}
