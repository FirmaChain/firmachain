package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TokenDataList: []TokenData{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in tokenData
	tokenDataIndexMap := make(map[string]struct{})

	for _, elem := range gs.TokenDataList {
		index := string(TokenDataKey(elem.TokenId))
		if _, ok := tokenDataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tokenData")
		}
		tokenDataIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
