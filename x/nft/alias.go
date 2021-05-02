package nft

import (
	"github.com/firmachain/FirmaChain/x/nft/keeper"
	"github.com/firmachain/FirmaChain/x/nft/types"
)

const (
	QueryNFT          = types.QueryNFT
	ModuleName        = types.ModuleName
	QuerierRoute      = types.QuerierRoute
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace

	EventTypeTransfer     = types.EventTypeTransfer
	AttributeKeyHash      = types.AttributeKeyHash
	AttributeKeyRecipient = types.AttributeKeyRecipient
	AttributeKeySender    = types.AttributeKeySender
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec

	NewNFT                 = types.NewNFT()
	NewMsgMintNFT          = types.NewMsgMintNFT
	NewMsgBurnNFT          = types.NewMsgBurnNFT
	NewMsgTransferNFT      = types.NewMsgTransferNFT
	NewMsgMultiTransferNFT = types.NewMsgMultiTransferNFT

	DefaultCodespace = types.DefaultCodespace

	ErrTokenNotFound = types.ErrTokenNotFound
	ErrInvalidHash   = types.ErrInvalidHash
	ErrExistedHash   = types.ErrExistedHash
	ErrNotOwnerToken = types.ErrNotOwnerToken
)

type (
	Keeper = keeper.Keeper

	NFT                 = types.NFT
	MsgMintNFT          = types.MsgMintNFT
	MsgBurnNFT          = types.MsgBurnNFT
	MsgTransferNFT      = types.MsgTransferNFT
	MsgMultiTransferNFT = types.MsgMultiTransferNFT
	QueryResNFT         = types.QueryResNFT
	Output              = types.Output
)
