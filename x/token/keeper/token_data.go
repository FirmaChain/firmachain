package keeper

import (
	"encoding/binary"
	"fmt"
	"regexp"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

const (
	maxSymbolLength = 20
	maxNameLength   = 40
	maxTokenValue   = 10000000000000000
	ufctTokenName   = "ufct"
)

func (k Keeper) CheckCommonError(tokenID string, symbol string, name string, totalSupply uint64) error {
	if len(symbol) > maxSymbolLength {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "input Symbol name lengh cannot exceed 20.")
	}

	if len(name) > maxNameLength {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "input Token name lengh cannot exceed 40.")
	}

	if tokenID == ufctTokenName {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "ufct is not supported.")
	}

	re := regexp.MustCompile("^[a-zA-Z0-9_-]*$")
	if !re.MatchString(tokenID) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "only alphabet, underscore, numbers are allowed for tokenID.")
	}

	if totalSupply > maxTokenValue {
		errStr := fmt.Sprintf("TotalSupply cannot exceed  %d", maxTokenValue)
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, errStr)
	}

	if totalSupply <= 0 {
		errStr := "TotalSupply cannot be minus or zero"
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, errStr)
	}

	return nil
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

// SetTokenData set a specific tokenData in the store from its index
func (k Keeper) SetTokenData(ctx sdk.Context, tokenData types.TokenData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenDataKeyPrefix))
	b := k.cdc.MustMarshal(&tokenData)
	store.Set(types.TokenDataKey(tokenData.TokenID), b)

	k.SetTokenDataToAccount(ctx, tokenData)
}

func (k Keeper) SetTokenDataToAccount(ctx sdk.Context, tokenData types.TokenData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerOfTokenKey))
	accountStore := prefix.NewStore(store, []byte(tokenData.Owner))

	accountStore.Set([]byte(tokenData.TokenID), GetBytesFromUInt64(1))
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

// GetAllTokenData returns all tokenData
func (k Keeper) GetAllTokenData(ctx sdk.Context) (list []types.TokenData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenDataKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokenData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
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
