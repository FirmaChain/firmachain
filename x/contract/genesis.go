package contract

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/v05/x/contract/keeper"
	"github.com/firmachain/firmachain/v05/x/contract/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the contractFile
	for _, elem := range genState.ContractFileList {
		k.SetContractFile(ctx, *elem)
	}

	// Set all the contractLog
	for _, elem := range genState.ContractLogList {
		// INFO: the hashStore value can be restored with ContractLog, so it is not stored separately.
		k.SetContractLog(ctx, *elem)
	}

	// Set contractLog count
	k.SetContractLogCount(ctx, genState.ContractLogCount)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// Get all contractFile
	contractFileList := k.GetAllContractFile(ctx)
	for _, elem := range contractFileList {
		elem := elem
		genesis.ContractFileList = append(genesis.ContractFileList, &elem)
	}

	// Get all contractLog
	contractLogList := k.GetAllContractLog(ctx)
	for _, elem := range contractLogList {
		elem := elem
		genesis.ContractLogList = append(genesis.ContractLogList, &elem)
	}

	// Set the current count
	genesis.ContractLogCount = k.GetContractLogCount(ctx)

	return genesis
}
