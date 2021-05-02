package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgMintNFT{}, "nft/MsgMintNFT", nil)
	cdc.RegisterConcrete(MsgTransferNFT{}, "nft/MsgTransferNFT", nil)
	cdc.RegisterConcrete(MsgMultiTransferNFT{}, "nft/MsgMultiTransferNFT", nil)
	cdc.RegisterConcrete(MsgBurnNFT{}, "nft/MsgBurnNFT", nil)
}
