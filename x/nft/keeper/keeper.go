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

func (k Keeper) IsDuplicateOwner(nft types.NFT, owner sdk.AccAddress) bool {
	return owner.Equals(nft.Owner)
}

func (k Keeper) GetNFT(ctx sdk.Context, hash string) types.NFT {
	store := ctx.KVStore(k.storeKey)

	if !k.IsTokenExisted(ctx, hash) {
		return types.NewNFT()
	}

	bz := store.Get([]byte(hash))

	var nft types.NFT
	k.cdc.MustUnmarshalBinaryBare(bz, &nft)

	return nft
}

func (k Keeper) InitNFT(ctx sdk.Context, hash string, tokenURI string, owner sdk.AccAddress) {
	nft := k.GetNFT(ctx, hash)
	nft.Hash = hash
	nft.TokenURI = tokenURI
	nft.Owner = owner

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(hash), k.cdc.MustMarshalBinaryBare(nft))
}

func (k Keeper) Mint(ctx sdk.Context, hash string, tokenURI string, owner sdk.AccAddress) error {
	if k.IsTokenExisted(ctx, hash) {
		return types.ErrExistedHash
	}

	nft := k.GetNFT(ctx, hash)

	if len(nft.Hash) == 0 {
		nft.Hash = hash
	}

	if len(nft.TokenURI) == 0 {
		nft.TokenURI = tokenURI
	}

	nft.Owner = owner

	k.SetNFT(ctx, hash, nft)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyHash, hash),
			sdk.NewAttribute(types.AttributeKeySender, owner.String()),
		),
	)

	return nil
}

func (k Keeper) Burn(ctx sdk.Context, hash string, owner sdk.AccAddress) error {
	if !k.IsTokenExisted(ctx, hash) {
		return types.ErrTokenNotFound
	}

	if owner.Empty() {
		return sdkerrors.ErrInvalidAddress
	}

	nft := k.GetNFT(ctx, hash)

	if !owner.Equals(nft.Owner) {
		return types.ErrNotOwnerToken
	}

	k.DeleteNFT(ctx, hash)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBurn,
			sdk.NewAttribute(types.AttributeKeyHash, hash),
			sdk.NewAttribute(types.AttributeKeySender, owner.String()),
		),
	)

	return nil
}

func (k Keeper) Transfer(ctx sdk.Context, hash string, owner sdk.AccAddress, recipient sdk.AccAddress) error {
	if !k.IsTokenExisted(ctx, hash) {
		return types.ErrTokenNotFound
	}

	if owner.Empty() || recipient.Empty() {
		return sdkerrors.ErrInvalidAddress
	}

	nft := k.GetNFT(ctx, hash)

	if !owner.Equals(nft.Owner) {
		return types.ErrNotOwnerToken
	}

	nft.Owner = recipient

	k.SetNFT(ctx, hash, nft)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransfer,
			sdk.NewAttribute(types.AttributeKeyHash, hash),
			sdk.NewAttribute(types.AttributeKeySender, owner.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
		),
	)

	return nil
}

func (k Keeper) MultiTransfer(ctx sdk.Context, owner sdk.AccAddress, outputs []types.Output) error {
	// Safety check ensuring that when sending coins the k must maintain the
	// Check supply invariant and validity of NFT.

	for _, out := range outputs {
		if err := out.ValidateBasic(); err != nil {
			return err
		}

		err := k.Transfer(ctx, out.Hash, owner, out.Recipient)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) SetNFT(ctx sdk.Context, hash string, nft types.NFT) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(hash), k.cdc.MustMarshalBinaryBare(nft))
}

func (k Keeper) DeleteNFT(ctx sdk.Context, hash string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(hash))
}

func (k Keeper) GetNFTsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
