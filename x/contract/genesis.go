package contract

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	ContractRecords []Contract `json:"contract_records"`
}

func NewGenesisState() GenesisState {

	return GenesisState{
		ContractRecords: nil,
	}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ContractRecords {
		if len(record.Owners) == 0 {
			return fmt.Errorf("invalid ContractRecords: Value: %v. Error: Missing Owner", record.Owners)
		}

		if record.Path == "" {
			return fmt.Errorf("invalid ContractRecords: Value: %s. Error: Missing Path", record.Path)
		}

		if record.Hash == "" {
			return fmt.Errorf("invalid ContractRecords: Value: %s. Error: Missing Hash", record.Hash)
		}

	}
	return nil
}

func DefaultGenesisState() GenesisState {

	return GenesisState{
		ContractRecords: []Contract{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.ContractRecords {
		keeper.InitContract(ctx, record.Hash, record.Path, record.Owners)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Contract
	iterator := k.GetContractsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		hash := string(iterator.Key())
		whois := k.GetContract(ctx, hash)
		records = append(records, whois)
	}

	return GenesisState{ContractRecords: records}
}
