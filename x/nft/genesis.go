package nft

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/firmachain/FirmaChain/x/nft/types"
)

func NewGenesisState() types.GenesisState {

	return types.GenesisState{
		NFTRecords: nil,
	}
}

func ValidateGenesis(data types.GenesisState) error {
	for _, record := range data.NFTRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid NFTRecords: Value: %v. Error: Missing Owner", record.Owner)
		}

		if record.Creator == nil {
			return fmt.Errorf("invalid NFTRecords: Value: %v. Error: Missing Creator", record.Creator)
		}

		if record.Description == "" {
			return fmt.Errorf("invalid NFTRecords: Value: %v. Error: Missing Description", record.Description)
		}

		if record.Image == "" {
			return fmt.Errorf("invalid NFTRecords: Value: %v. Error: Missing Image", record.Image)
		}

		if record.Hash == "" {
			return fmt.Errorf("invalid NFTRecords: Value: %s. Error: Missing Hash", record.Hash)
		}

		if record.TokenURI == "" {
			return fmt.Errorf("invalid NFTRecords: Value: %s. Error: Missing TokenURI", record.TokenURI)
		}

	}
	return nil
}

func DefaultGenesisState() types.GenesisState {

	return types.GenesisState{
		NFTRecords: []NFT{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data types.GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.NFTRecords {
		keeper.InitNFT(ctx, record.Hash, record.TokenURI, record.Owner, record.Description, record.Image)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) types.GenesisState {
	var records []NFT
	iterator := k.GetNFTsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		hash := string(iterator.Key())
		whois := k.GetNFT(ctx, hash)
		records = append(records, whois)
	}

	return types.GenesisState{NFTRecords: records}
}
