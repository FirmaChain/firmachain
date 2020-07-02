package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/FirmaChain/x/contract/types"
)

type Keeper struct {
	cdc      *codec.Codec
	storeKey sdk.StoreKey
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

func (k Keeper) IsContractPresent(ctx sdk.Context, hash string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(hash))
}

func (k Keeper) GetContract(ctx sdk.Context, hash string) types.Contract {
	store := ctx.KVStore(k.storeKey)
	if !k.IsContractPresent(ctx, hash) {
		return types.NewContract()
	}
	bz := store.Get([]byte(hash))

	var contract types.Contract
	k.cdc.MustUnmarshalBinaryBare(bz, &contract)
	return contract
}

func (k Keeper) SetContract(ctx sdk.Context, path string, hash string, owner sdk.AccAddress) {
	contract := k.GetContract(ctx, hash)
	contract.Hash = hash
	contract.Path = path
	contract.Owner = owner
	k.AddContract(ctx, hash, contract)
}

func (k Keeper) AddContract(ctx sdk.Context, hash string, contract types.Contract) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(hash), k.cdc.MustMarshalBinaryBare(contract))
}

func (k Keeper) GetContractsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
