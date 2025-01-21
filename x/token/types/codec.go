package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgCreateToken{}, "token/CreateToken")
	legacy.RegisterAminoMsg(cdc, &MsgMint{}, "token/Mint")
	legacy.RegisterAminoMsg(cdc, &MsgBurn{}, "token/Burn")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateTokenURI{}, "token/UpdateTokenURI")
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateToken{},
		&MsgMint{},
		&MsgBurn{},
		&MsgUpdateTokenURI{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
