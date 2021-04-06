package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/FirmaChain/x/nft/types"
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

func (k Keeper) IsTokenExisted(ctx sdk.Context, hash string) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has([]byte(hash))
}

func (k Keeper) IsDuplicateOwner(nft types.NFToken, owner sdk.AccAddress) bool {
	return owner.Equals(nft.Owner)
}

func (k Keeper) GetNFToken(ctx sdk.Context, hash string) types.NFToken {
	store := ctx.KVStore(k.storeKey)

	if !k.IsTokenExisted(ctx, hash) {
		return types.NewNFToken()
	}

	bz := store.Get([]byte(hash))

	var nft types.NFToken
	k.cdc.MustUnmarshalBinaryBare(bz, &nft)

	return nft
}

func (k Keeper) InitNFToken(ctx sdk.Context, hash string, tokenURI string, owner sdk.AccAddress) {
	nft := k.GetNFToken(ctx, hash)
	nft.Hash = hash
	nft.TokenURI = tokenURI
	nft.Owner = owner

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(hash), k.cdc.MustMarshalBinaryBare(nft))
}

func (k Keeper) AddNFToken(ctx sdk.Context, hash string, tokenURI string, owner sdk.AccAddress) error {
	if k.IsTokenExisted(ctx, hash) {
		return types.ErrNFTokenDuplicated
	}

	nft := k.GetNFToken(ctx, hash)

	if len(nft.Hash) == 0 {
		nft.Hash = hash
	}

	if len(nft.TokenURI) == 0 {
		nft.TokenURI = tokenURI
	}

	nft.Owner = owner

	k.SetNFToken(ctx, hash, nft)

	return nil
}

func (k Keeper) TransferNFToken(ctx sdk.Context, hash string, recipient sdk.AccAddress) error {
	if !k.IsTokenExisted(ctx, hash) {
		return types.ErrTokenDoesNotExist
	}

	if recipient == nil {
		return sdkerrors.ErrInvalidAddress
	}

	nft := k.GetNFToken(ctx, hash)
	nft.Owner = recipient

	k.SetNFToken(ctx, hash, nft)

	return nil
}

func (k Keeper) SetNFToken(ctx sdk.Context, hash string, nft types.NFToken) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(hash), k.cdc.MustMarshalBinaryBare(nft))
}

func (k Keeper) GetNFTokensIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
