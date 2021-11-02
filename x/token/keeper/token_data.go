package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/x/token/types"
)

// SetTokenData set a specific tokenData in the store from its index
func (k Keeper) SetTokenData(ctx sdk.Context, tokenData types.TokenData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenDataKeyPrefix))
	b := k.cdc.MustMarshal(&tokenData)
	store.Set(types.TokenDataKey(
		tokenData.TokenID), b)
}

func (k Keeper) AddTokenDataToAccount(ctx sdk.Context, address string, tokenID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenDataAccountMapKey))
	accountStore := prefix.NewStore(store, []byte(address))

	accountStore.Set([]byte(tokenID), []byte(tokenID))
}

// GetTokenData returns a tokenData from its index
func (k Keeper) GetTokenData(
	ctx sdk.Context,
	tokenID string,

) (val types.TokenData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenDataKeyPrefix))

	b := store.Get(types.TokenDataKey(
		tokenID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTokenData removes a tokenData from the store
func (k Keeper) RemoveTokenData(
	ctx sdk.Context,
	tokenID string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenDataKeyPrefix))
	store.Delete(types.TokenDataKey(
		tokenID,
	))
}

// GetAllTokenData returns all tokenData
func (k Keeper) GetAllTokenData(ctx sdk.Context) (list []types.TokenData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokenData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
