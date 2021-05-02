package types

type GenesisState struct {
	NFTRecords []NFT `json:"nft_records"`
}

func NewGenesisState() GenesisState {

	return GenesisState{
		NFTRecords: nil,
	}
}
