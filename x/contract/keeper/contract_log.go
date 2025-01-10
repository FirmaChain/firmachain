package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/v05/x/contract/types"
)

func (k Keeper) AppendContractLog(ctx sdk.Context, contractLog types.ContractLog) uint64 {
	count := k.GetContractLogCount(ctx)
	contractLog.Id = count

	k.SetContractLog(ctx, contractLog)
	k.SetContractLogCount(ctx, count+1)

	return count
}

func (k Keeper) SetContractLogToHashStore(ctx sdk.Context, contractLog types.ContractLog) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogHashKey))

	hashStore := prefix.NewStore(store, []byte(contractLog.ContractHash))
	hashStore.Set(GetBytesFromUInt64(contractLog.Id), GetBytesFromUInt64(1))
}

func (k Keeper) SetContractLogCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogDataTotalKey))
	byteKey := types.KeyPrefix(types.ContractLogDataTotalKey)
	bz := GetBytesFromUInt64(count)
	store.Set(byteKey, bz)
}

func (k Keeper) SetContractLog(ctx sdk.Context, contractLog types.ContractLog) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogDataKey))
	b := k.cdc.MustMarshal(&contractLog)
	store.Set(GetBytesFromUInt64(contractLog.Id), b)

	k.SetContractLogToHashStore(ctx, contractLog)
}

func (k Keeper) GetContractLog(ctx sdk.Context, id uint64) types.ContractLog {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogDataKey))
	var contractLog types.ContractLog
	k.cdc.MustUnmarshal(store.Get(GetBytesFromUInt64(id)), &contractLog)
	return contractLog
}

// GetContractLogCount get the total number of TypeName.LowerCamel
func (k Keeper) GetContractLogCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogDataTotalKey))
	byteKey := types.KeyPrefix(types.ContractLogDataTotalKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count := GetUInt64FromBytes(bz)

	return count
}

func (k Keeper) GetContractLogOwner(ctx sdk.Context, id uint64) string {
	return k.GetContractLog(ctx, id).Creator
}

func (k Keeper) GetAllContractLog(ctx sdk.Context) (list []types.ContractLog) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogDataKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ContractLog
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) HasContractLog(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogDataKey))
	return store.Has(GetBytesFromUInt64(id))
}

// GetBytesFromUInt64 returns the byte representation of the Uint64
func GetBytesFromUInt64(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetUInt64FromBytes returns uint64 format from a byte array
func GetUInt64FromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
