package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/v05/x/nft/types"
)

func (k Keeper) AppendNftItem(
	ctx sdk.Context,
	nftItem types.NftItem,
) uint64 {

	// Create the nftItem
	count := k.GetNftItemCount(ctx)

	// Set the ID of the appended value
	nftItem.Id = count

	k.SetNftItem(ctx, nftItem)
	k.SetNftItemCount(ctx, count+1)

	return count
}

func (k Keeper) RemoveNftItemToAccount(ctx sdk.Context, address string, nftID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerOfNftKey))
	accountStore := prefix.NewStore(store, []byte(address))

	accountStore.Delete(GetBytesFromUInt64(nftID))

	total := k.GetNftTotalToAccount(ctx, address)
	k.SetNftTotalToAccount(ctx, address, total-1)
}

// RemoveNftItem removes a nftItem from the store
func (k Keeper) RemoveNftItem(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemDataKey))
	store.Delete(GetBytesFromUInt64(id))
}

func (k Keeper) SetNftItemCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemTotalKey))

	byteKey := types.KeyPrefix(types.NftItemTotalKey)
	bz := []byte(strconv.FormatUint(count, 10))

	store.Set(byteKey, bz)
}

// SetNftItem set a specific nftItem in the store
func (k Keeper) SetNftItem(ctx sdk.Context, nftItem types.NftItem) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemDataKey))
	b := k.cdc.MustMarshal(&nftItem)
	store.Set(GetBytesFromUInt64(nftItem.Id), b)

	k.SetNftItemToAccount(ctx, nftItem.Owner, nftItem.Id)
}

func (k Keeper) SetNftItemToAccount(ctx sdk.Context, address string, nftID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerOfNftKey))
	accountStore := prefix.NewStore(store, []byte(address))

	accountStore.Set(GetBytesFromUInt64(nftID), GetBytesFromUInt64(1))

	total := k.GetNftTotalToAccount(ctx, address)
	k.SetNftTotalToAccount(ctx, address, total+1)
}

func (k Keeper) SetNftTotalToAccount(ctx sdk.Context, address string, total uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerOfNftTotalKey))
	accountStore := prefix.NewStore(store, []byte(address))

	byteKey := types.KeyPrefix(types.OwnerOfNftTotalKey)
	accountStore.Set(byteKey, GetBytesFromUInt64(total))
}

// GetNftItemCount get the total number of TypeName.LowerCamel
func (k Keeper) GetNftItemCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemTotalKey))
	byteKey := types.KeyPrefix(types.NftItemTotalKey)
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

func (k Keeper) GetNftTotalToAccount(ctx sdk.Context, address string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerOfNftTotalKey))
	accountStore := prefix.NewStore(store, []byte(address))

	byteKey := types.KeyPrefix(types.OwnerOfNftTotalKey)
	byteTotal := accountStore.Get(byteKey)

	// Count doesn't exist: no element
	if byteTotal == nil {
		return 0
	}

	// Parse bytes
	count := GetUInt64FromBytes(byteTotal)

	return count
}

// GetNftItem returns a nftItem from its id
func (k Keeper) GetNftItem(ctx sdk.Context, id uint64) types.NftItem {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemDataKey))
	var nftItem types.NftItem
	k.cdc.MustUnmarshal(store.Get(GetBytesFromUInt64(id)), &nftItem)
	return nftItem
}

// GetNftItemOwner returns the creator of the
func (k Keeper) GetNftItemOwner(ctx sdk.Context, id uint64) string {
	return k.GetNftItem(ctx, id).Owner
}

// GetAllNftItem returns all nftItem
func (k Keeper) GetAllNftItem(ctx sdk.Context) (list []types.NftItem) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemDataKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftItem
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// HasNftItem checks if the nftItem exists in the store
func (k Keeper) HasNftItem(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemDataKey))
	return store.Has(GetBytesFromUInt64(id))
}

// GetBytesFromUInt64 returns the byte representation of the UInt64
func GetBytesFromUInt64(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetUInt64FromBytes returns uint64 format from a byte array
func GetUInt64FromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
