package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	// Set all the tokenData
	for _, elem := range genState.TokenDataList {
		k.SetTokenData(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.TokenDataList = k.GetAllTokenData(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
