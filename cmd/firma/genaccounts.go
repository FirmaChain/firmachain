package main

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
)

const (
	flagClientHome = "home-client"
)

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
func AddGenesisAccountCmd(ctx *server.Context, cdc *codec.Codec,
	defaultNodeHome, defaultClientHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-account [address_or_key_name] [coin][,[coin]]",
		Short: "Add genesis account to genesis.json",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				inBuf := bufio.NewReader(cmd.InOrStdin())

				// attempt to lookup address from Keybase if no address was provided
				kb, err := keys.NewKeyring(
					sdk.KeyringServiceName(),
					viper.GetString(flags.FlagKeyringBackend),
					viper.GetString(flagClientHome),
					inBuf,
				)
				if err != nil {
					return err
				}

				info, err := kb.Get(args[0])
				if err != nil {
					return fmt.Errorf("failed to get address from Keybase: %w", err)
				}

				addr = info.GetAddress()
			}

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return fmt.Errorf("failed to parse coins: %w", err)
			}

			var genAccount authexported.GenesisAccount

			baseAccount := types.NewBaseAccountWithAddress(addr)
			baseAccount.Coins = coins

			genAccount = &baseAccount
			if err := genAccount.Validate(); err != nil {
				return err
			}

			genFile := config.GenesisFile()
			appState, genDoc, err := genutil.GenesisStateFromGenFile(cdc, genFile)
			if err != nil {
				return err
			}

			var genesisState auth.GenesisState
			cdc.MustUnmarshalJSON(appState[auth.ModuleName], &genesisState)

			if genesisState.Accounts.Contains(addr) {
				return fmt.Errorf("cannot add account at existing address %s", addr)
			}

			genesisState.Accounts = append(genesisState.Accounts, genAccount)

			genesisStateBz := cdc.MustMarshalJSON(genesisState)
			appState[auth.ModuleName] = genesisStateBz

			appStateJSON, err := cdc.MarshalJSON(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			// export app state
			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultClientHome, "client's home directory")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test)")
	return cmd
}
