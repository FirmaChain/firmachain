package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NftItemList: []*NftItem{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in nftItem
	nftItemIdMap := make(map[uint64]bool)

	for _, elem := range gs.NftItemList {
		if _, ok := nftItemIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for nftItem")
		}
		nftItemIdMap[elem.Id] = true
	}

	return nil
}
