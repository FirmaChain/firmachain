package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/firmachain/firmachain/x/contract/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group contract queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdIsContractOwner())
	cmd.AddCommand(CmdListContractFile())
	cmd.AddCommand(CmdShowContractFile())
	cmd.AddCommand(CmdListContractLog())
	cmd.AddCommand(CmdShowContractLog())
	cmd.AddCommand(CmdContractListFromHash())

	return cmd
}
