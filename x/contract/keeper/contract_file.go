package keeper

import (
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/v05/x/contract/types"
)

// SetContractFile set a specific contractFile in the store from its index
func (k Keeper) SetContractFile(ctx sdk.Context, contractFile types.ContractFile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractFileKey))
	//b := k.cdc.MustMarshalBinaryBare(&contractFile)
	b := k.cdc.MustMarshal(&contractFile)
	store.Set(types.KeyPrefix(contractFile.FileHash), b)
}

// GetContractFile returns a contractFile from its index
func (k Keeper) GetContractFile(ctx sdk.Context, index string) (val types.ContractFile, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractFileKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContractFile removes a contractFile from the store
func (k Keeper) RemoveContractFile(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractFileKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllContractFile returns all contractFile
func (k Keeper) GetAllContractFile(ctx sdk.Context) (list []types.ContractFile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractFileKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ContractFile
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
