package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgMint{}, "nft/MsgMint", nil)
	cdc.RegisterConcrete(MsgTransfer{}, "nft/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgMultiTransfer{}, "nft/MsgMultiTransfer", nil)
	cdc.RegisterConcrete(MsgBurn{}, "nft/MsgBurn", nil)
}
