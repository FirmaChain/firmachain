package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/x/nft/types"
)

// GetNftItemCount get the total number of TypeName.LowerCamel
func (k Keeper) GetNftItemCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemCountKey))
	byteKey := types.KeyPrefix(types.NftItemCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to uint64
		panic("cannot decode count")
	}

	return count
}

func (k Keeper) SetNftItemCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemCountKey))

	byteKey := types.KeyPrefix(types.NftItemCountKey)
	bz := []byte(strconv.FormatUint(count, 10))

	store.Set(byteKey, bz)
}

func (k Keeper) AddNftItemToAccount(ctx sdk.Context, address string, nftID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemAccountMapKey))
	accountStore := prefix.NewStore(store, []byte(address))

	accountStore.Set(GetNftItemIDBytes(nftID), GetNftItemIDBytes(1))
}

func (k Keeper) RemoveNftItemToAccount(ctx sdk.Context, address string, nftID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemAccountMapKey))
	accountStore := prefix.NewStore(store, []byte(address))

	accountStore.Delete(GetNftItemIDBytes(nftID))
}

func (k Keeper) AppendNftItem(
	ctx sdk.Context,
	nftItem types.NftItem,
) uint64 {

	// Create the nftItem
	count := k.GetNftItemCount(ctx)

	// Set the ID of the appended value
	nftItem.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&nftItem)
	store.Set(GetNftItemIDBytes(nftItem.Id), appendedValue)

	// Update nftItem count
	k.SetNftItemCount(ctx, count+1)

	return count
}

// SetNftItem set a specific nftItem in the store
func (k Keeper) SetNftItem(ctx sdk.Context, nftItem types.NftItem) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemKey))
	b := k.cdc.MustMarshalBinaryBare(&nftItem)
	store.Set(GetNftItemIDBytes(nftItem.Id), b)
}

// GetNftItem returns a nftItem from its id
func (k Keeper) GetNftItem(ctx sdk.Context, id uint64) types.NftItem {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemKey))
	var nftItem types.NftItem
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetNftItemIDBytes(id)), &nftItem)
	return nftItem
}

// HasNftItem checks if the nftItem exists in the store
func (k Keeper) HasNftItem(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemKey))
	return store.Has(GetNftItemIDBytes(id))
}

// GetNftItemOwner returns the creator of the
func (k Keeper) GetNftItemOwner(ctx sdk.Context, id uint64) string {
	return k.GetNftItem(ctx, id).Owner
}

// RemoveNftItem removes a nftItem from the store
func (k Keeper) RemoveNftItem(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemKey))
	store.Delete(GetNftItemIDBytes(id))
}

// GetAllNftItem returns all nftItem
func (k Keeper) GetAllNftItem(ctx sdk.Context) (list []types.NftItem) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftItem
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetNftItemIDBytes returns the byte representation of the ID
func GetNftItemIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetNftItemIDFromBytes returns ID in uint64 format from a byte array
func GetNftItemIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
