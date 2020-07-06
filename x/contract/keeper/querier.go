package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		return QueryContract(ctx, path, req, keeper)
	}
}

func QueryContract(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	contract := keeper.GetContract(ctx, path[0])

	res, err := codec.MarshalJSONIndent(keeper.cdc, contract)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
