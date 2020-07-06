package contract

import (
	"github.com/firmachain/FirmaChain/x/contract/keeper"
	"github.com/firmachain/FirmaChain/x/contract/types"
)

const (
	ModuleName    = types.ModuleName
	RouterKey     = types.RouterKey
	StoreKey      = types.StoreKey
	QuerierRoute  = types.QuerierRoute
	QueryContract = types.QueryContract
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec

	NewContract       = types.NewContract
	NewMsgAddContract = types.NewMsgAddContract

	DefaultCodespace = types.DefaultCodespace

	ErrContractDoesNotExist = types.ErrContractDoesNotExist
	ErrContractInvalid      = types.ErrContractInvalid
	ErrContractDuplicated   = types.ErrContractDuplicated
)

type (
	Keeper = keeper.Keeper

	Contract         = types.Contract
	MsgAddContract   = types.MsgAddContract
	QueryResContract = types.QueryResContract
)
