package contract

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/firmachain/FirmaChain/x/contract/types"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/firmachain/FirmaChain/x/contract/client/cli"
	"github.com/firmachain/FirmaChain/x/contract/client/rest"
	"github.com/firmachain/FirmaChain/x/contract/simulation"
)

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// app module Basics object
type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

// Validation check of the Genesis
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data types.GenesisState
	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	// Once json successfully marshalled, passes along to genesis.go
	return ValidateGenesis(data)
}

// Register rest routes
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr)
}

// Get the root query command of this module
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(cdc)
}

// Get the root tx command of this module
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(cdc)
}

type AppModule struct {
	AppModuleBasic
	keeper        Keeper
	accountKeeper auth.AccountKeeper
}

func (am AppModule) GenerateGenesisState(input *module.SimulationState) {
	simulation.RandomizedGenState(input)
}

func (am AppModule) ProposalContents(simState module.SimulationState) []sim.WeightedProposalContent {
	return nil
}

func (am AppModule) RandomizedParams(r *rand.Rand) []sim.ParamChange {
	return simulation.ParamChanges(r)
}

func (am AppModule) RegisterStoreDecoder(registry sdk.StoreDecoderRegistry) {}

func (am AppModule) WeightedOperations(simState module.SimulationState) []sim.WeightedOperation {
	return simulation.WeightedOperations(simState.AppParams, simState.Cdc, am.accountKeeper, am.keeper)
}

// NewAppModule creates a new AppModule Object
func NewAppModule(k Keeper, accountKeeper auth.AccountKeeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
		accountKeeper:  accountKeeper,
	}
}

func (AppModule) Name() string {
	return ModuleName
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() string {
	return RouterKey
}

func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}
func (am AppModule) QuerierRoute() string {
	return ModuleName
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (am AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	return InitGenesis(ctx, am.keeper, genesisState)
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return ModuleCdc.MustMarshalJSON(gs)
}
