package auth

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the auth module.
type AppModuleBasic struct{}

// Name returns the auth module's name
func (AppModuleBasic) Name() string {
	return auth.AppModuleBasic{}.Name()
}

// RegisterCodec registers the auth module's types for the given codec.
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	auth.RegisterCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the auth
// module.
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return auth.AppModuleBasic{}.DefaultGenesis()
}

// ValidateGenesis performs genesis state validation for the auth module.
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	return auth.AppModuleBasic{}.ValidateGenesis(bz)
}

// RegisterRESTRoutes registers the REST routes for the auth module.
func (AppModuleBasic) RegisterRESTRoutes(cliCtx context.CLIContext, route *mux.Router) {
	auth.AppModuleBasic{}.RegisterRESTRoutes(cliCtx, route)
}

// GetTxCmd returns the root tx command for the auth module.
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return auth.AppModuleBasic{}.GetTxCmd(cdc)
}

// GetQueryCmd returns the root query command for the auth module.
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return auth.AppModuleBasic{}.GetQueryCmd(cdc)
}

type AppModule struct {
	AppModuleBasic
	accountKeeper auth.AccountKeeper
}

// NewAppModule creates a new AppModule Object
func NewAppModule(accountKeeper auth.AccountKeeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		accountKeeper:  accountKeeper,
	}
}

func (AppModule) Name() string {
	return auth.AppModule{}.Name()
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() string {
	return auth.AppModule{}.Route()
}

func (am AppModule) NewHandler() sdk.Handler {
	return auth.AppModule{}.NewHandler()
}
func (am AppModule) QuerierRoute() string {
	return auth.AppModule{}.QuerierRoute()
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
	return auth.AppModule{}.NewQuerierHandler()
}

func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (am AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	return auth.AppModule{}.InitGenesis(ctx, data)
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return auth.AppModule{}.ExportGenesis(ctx)
}
