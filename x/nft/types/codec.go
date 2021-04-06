package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddNFToken{}, "nft/AddNFToken", nil)
	cdc.RegisterConcrete(MsgTransferNFToken{}, "nft/TransferNFToken", nil)
}
