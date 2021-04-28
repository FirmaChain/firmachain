package nft

import (
	"github.com/firmachain/FirmaChain/x/nft/keeper"
	"github.com/firmachain/FirmaChain/x/nft/types"
)

const (
	QueryNFToken      = types.QueryNFToken
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

	NewNFToken          = types.NewNFToken()
	NewMsgMint          = types.NewMsgMint
	NewMsgBurn          = types.NewMsgBurn
	NewMsgTransfer      = types.NewMsgTransfer
	NewMsgMultiTransfer = types.NewMsgMultiTransfer

	DefaultCodespace = types.DefaultCodespace

	ErrTokenNotFound = types.ErrTokenNotFound
	ErrInvalidHash   = types.ErrInvalidHash
	ErrExistedHash   = types.ErrExistedHash
	ErrNotOwnerToken = types.ErrNotOwnerToken
)

type (
	Keeper = keeper.Keeper

	NFToken          = types.NFToken
	MsgMint          = types.MsgMint
	MsgBurn          = types.MsgBurn
	MsgTransfer      = types.MsgTransfer
	MsgMultiTransfer = types.MsgMultiTransfer
	QueryResNFToken  = types.QueryResNFToken
	Output           = types.Output
)
