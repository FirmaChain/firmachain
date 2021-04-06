package types

type GenesisState struct {
	NFTokenRecords []NFToken `json:"nft_records"`
}

func NewGenesisState() GenesisState {

	return GenesisState{
		NFTokenRecords: nil,
	}
}
