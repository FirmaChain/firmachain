package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		return QueryContract(ctx, path[0], req, keeper)
	}
}

func QueryContract(ctx sdk.Context, hash string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	contract := keeper.GetNFT(ctx, hash)

	res, err := codec.MarshalJSONIndent(keeper.cdc, contract)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
