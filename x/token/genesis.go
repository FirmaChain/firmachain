package token

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/v05/x/token/keeper"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the tokenData
	for _, elem := range genState.TokenDataList {
		k.SetTokenData(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.TokenDataList = k.GetAllTokenData(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
