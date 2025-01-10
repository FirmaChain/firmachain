package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/v05/x/nft/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	// Set all the nftItem
	for _, elem := range genState.NftItemList {

		// INFO: the nftList and totalCount for owner can be restored with NftItem, so it is not stored separately.
		k.SetNftItem(ctx, *elem)
	}

	// Set nftItem count
	k.SetNftItemCount(ctx, genState.NftItemCount)
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// Get all nftItem
	nftItemList := k.GetAllNftItem(ctx)
	for _, elem := range nftItemList {
		elem := elem
		genesis.NftItemList = append(genesis.NftItemList, &elem)
	}

	// Set the current count
	genesis.NftItemCount = k.GetNftItemCount(ctx)

	return genesis
}
