package nft

import (
	"github.com/firmachain/FirmaChain/x/nft/keeper"
	"github.com/firmachain/FirmaChain/x/nft/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace

	QueryNFToken = types.QueryNFToken
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec

	NewNFToken            = types.NewNFToken()
	NewMsgAddNFToken      = types.NewMsgAddNFToken
	NewMsgTransferNFToken = types.NewMsgTransferNFToken

	DefaultCodespace = types.DefaultCodespace

	ErrTokenNotFound = types.ErrTokenNotFound
	ErrInvalidHash   = types.ErrInvalidHash
	ErrExistedHash   = types.ErrExistedHash
	ErrNotOwnerToken = types.ErrNotOwnerToken
)

type (
	Keeper = keeper.Keeper

	NFToken            = types.NFToken
	MsgAddNFToken      = types.MsgAddNFToken
	MsgTransferNFToken = types.MsgTransferNFToken
	QueryResNFToken    = types.QueryResNFToken
)
