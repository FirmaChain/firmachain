package simulation

// DONTCOVER

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/firmachain/FirmaChain/x/contract/types"
)

// RandomizedGenState generates a random GenesisState for contract
func RandomizedGenState(simState *module.SimulationState) {
	contractGenesis := types.NewGenesisState()

	fmt.Printf("Selected randomly generated contract parameters:\n%s\n", codec.MustMarshalJSONIndent(simState.Cdc, contractGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(contractGenesis)
}
