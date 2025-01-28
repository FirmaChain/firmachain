package app

import (
	appparams "github.com/firmachain/firmachain/v05/app/params"

	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/codec"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/gogoproto/proto"
)

// NewInterfaceRegistry returns a new codectypes.InterfaceRegistry with standard options.
func NewInterfaceRegistry() (codectypes.InterfaceRegistry, error) {
	interfaceRegistry, err := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec:          addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
			ValidatorAddressCodec: addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		},
	})
	return interfaceRegistry, err
}

// MakeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func MakeEncodingConfig() appparams.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry, err := NewInterfaceRegistry()
	if err != nil {
		panic(err)
	}
	cdc := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(cdc, tx.DefaultSignModes)
	return appparams.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         cdc,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}
