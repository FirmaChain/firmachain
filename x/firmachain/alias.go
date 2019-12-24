package firmachain

import (
	"github.com/firmachain/FirmaChain/x/firmachain/internal/keeper"
	"github.com/firmachain/FirmaChain/x/firmachain/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey

	Denom = types.denom
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec

	NewContract       = types.NewContract
	NewMsgAddContract = types.NewMsgAddContract
)

type (
	Keeper = keeper.Keeper

	Contract         = types.Contract
	MsgAddContract   = types.MsgAddContract
	QueryResContract = types.QueryResContract
)
