package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/FirmaChain/x/contract/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case types.QueryContract:
			return QueryContract(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown contract query endpoint")
		}
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
