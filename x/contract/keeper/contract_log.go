package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/x/contract/types"
)

// GetContractLogCount get the total number of TypeName.LowerCamel
func (k Keeper) GetContractLogCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogCountKey))
	byteKey := types.KeyPrefix(types.ContractLogCountKey)
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

// SetContractLogCount set the total number of contractLog
func (k Keeper) SetContractLogCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogCountKey))
	byteKey := types.KeyPrefix(types.ContractLogCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendContractLog appends a contractLog in the store with a new id and update the count
func (k Keeper) AppendContractLog(
	ctx sdk.Context,
	contractLog types.ContractLog,
) uint64 {
	// Create the contractLog
	count := k.GetContractLogCount(ctx)

	// Set the ID of the appended value
	contractLog.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogKey))
	appendedValue := k.cdc.MustMarshal(&contractLog)
	store.Set(GetContractLogIDBytes(contractLog.Id), appendedValue)

	// Update contractLog count
	k.SetContractLogCount(ctx, count+1)

	return count
}

// SetContractLog set a specific contractLog in the store
func (k Keeper) SetContractLog(ctx sdk.Context, contractLog types.ContractLog) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogKey))
	b := k.cdc.MustMarshal(&contractLog)
	store.Set(GetContractLogIDBytes(contractLog.Id), b)
}

// GetContractLog returns a contractLog from its id
func (k Keeper) GetContractLog(ctx sdk.Context, id uint64) types.ContractLog {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogKey))
	var contractLog types.ContractLog
	k.cdc.MustUnmarshal(store.Get(GetContractLogIDBytes(id)), &contractLog)
	return contractLog
}

// HasContractLog checks if the contractLog exists in the store
func (k Keeper) HasContractLog(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogKey))
	return store.Has(GetContractLogIDBytes(id))
}

// GetContractLogOwner returns the creator of the
func (k Keeper) GetContractLogOwner(ctx sdk.Context, id uint64) string {
	return k.GetContractLog(ctx, id).Creator
}

// RemoveContractLog removes a contractLog from the store
func (k Keeper) RemoveContractLog(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogKey))
	store.Delete(GetContractLogIDBytes(id))
}

// GetAllContractLog returns all contractLog
func (k Keeper) GetAllContractLog(ctx sdk.Context) (list []types.ContractLog) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ContractLog
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetContractLogIDBytes returns the byte representation of the ID
func GetContractLogIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetContractLogIDFromBytes returns ID in uint64 format from a byte array
func GetContractLogIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
