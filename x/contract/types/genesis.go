package types

type GenesisState struct {
	ContractRecords []Contract `json:"contract_records"`
}

func NewGenesisState() GenesisState {

	return GenesisState{
		ContractRecords: nil,
	}
}
