package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ContractFileList: []*ContractFile{},
		ContractLogList:  []*ContractLog{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in contractFile
	contractFileIndexMap := make(map[string]bool)

	for _, elem := range gs.ContractFileList {
		if _, ok := contractFileIndexMap[elem.FileHash]; ok {
			return fmt.Errorf("duplicated index for contractFile")
		}
		contractFileIndexMap[elem.FileHash] = true
	}
	// Check for duplicated ID in contractLog
	contractLogIdMap := make(map[uint64]bool)

	for _, elem := range gs.ContractLogList {
		if _, ok := contractLogIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for contractLog")
		}
		contractLogIdMap[elem.Id] = true
	}

	return nil
}
