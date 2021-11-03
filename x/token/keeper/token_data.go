package keeper

import (
	"regexp"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/x/token/types"
)

const (
	maxSymbolLength = 20
	maxNameLength   = 40
	maxTokenValue   = 10000000000
	ufctTokenName   = "ufct"
)

func (k Keeper) CheckCommonError(tokenID string, symbol string, name string, totalSupply uint64) error {
	if len(symbol) > maxSymbolLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "input Symbol name lengh cannot exceed 20.")
	}

	if len(name) > maxNameLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "input Token name lengh cannot exceed 40.")
	}

	if tokenID == ufctTokenName {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "ufct is not supported.")
	}

	re := regexp.MustCompile("^[a-zA-Z0-9_-]*$")
	if !re.MatchString(tokenID) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "only alphabet, underscore, numbers are allowed for tokenID.")
	}

	if totalSupply > maxTokenValue {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "TotalSupply cannot exceed 10000000000")
	}

	return nil
}

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
