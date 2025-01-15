package app

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	//"strings"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	math "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	"cosmossdk.io/x/evidence"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	evidencetypes "cosmossdk.io/x/evidence/types"
	ctmos "github.com/cometbft/cometbft/libs/os"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/ibc-go/modules/capability"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	"github.com/firmachain/firmachain/v05/app/openapiconsole"

	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"

	//clienthelpers "cosmossdk.io/client/v2/helpers"
	"cosmossdk.io/x/feegrant"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"

	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	"cosmossdk.io/x/upgrade"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"

	ibcporttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	"github.com/spf13/cast"

	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	dbm "github.com/cosmos/cosmos-db"

	upgrades "github.com/firmachain/firmachain/v05/app/upgrades"
	v04 "github.com/firmachain/firmachain/v05/app/upgrades/v04"
	v05 "github.com/firmachain/firmachain/v05/app/upgrades/v05"
	"github.com/firmachain/firmachain/v05/client/docs"

	contractmodule "github.com/firmachain/firmachain/v05/x/contract"
	contractmodulekeeper "github.com/firmachain/firmachain/v05/x/contract/keeper"
	contractmoduletypes "github.com/firmachain/firmachain/v05/x/contract/types"
	nftmodule "github.com/firmachain/firmachain/v05/x/nft"
	nftmodulekeeper "github.com/firmachain/firmachain/v05/x/nft/keeper"
	nftmoduletypes "github.com/firmachain/firmachain/v05/x/nft/types"

	tokenmodule "github.com/firmachain/firmachain/v05/x/token"
	tokenmodulekeeper "github.com/firmachain/firmachain/v05/x/token/keeper"
	tokenmoduletypes "github.com/firmachain/firmachain/v05/x/token/types"

	"github.com/CosmWasm/wasmd/x/wasm"

	ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"

	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"

	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/codec/address"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"
	"github.com/cosmos/cosmos-sdk/std"
	sigtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
	authtxconfig "github.com/cosmos/cosmos-sdk/x/auth/tx/config"
	consensusparams "github.com/cosmos/cosmos-sdk/x/consensus"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	govv1beta "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/gogoproto/proto"
	packetforward "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward"
	packetforwardkeeper "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/keeper"
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"
	icq "github.com/cosmos/ibc-apps/modules/async-icq/v8"
	icqkeeper "github.com/cosmos/ibc-apps/modules/async-icq/v8/keeper"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v8/types"
	ibc_hooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8"
	ibchookskeeper "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/keeper"
	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/types"
	ibcfee "github.com/cosmos/ibc-go/v8/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/firmachain/firmachain/v05/app/keepers"
)

const (
	AccountAddressPrefix = "firma"
	Name                 = "firmachain"
	NodeDir              = ".firmachain"
)

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals
var (
// If EnabledSpecificProposals is "", and this is "true", then enable all x/wasm proposals.
// If EnabledSpecificProposals is "", and this is not "true", then disable all x/wasm proposals.
// ProposalsEnabled = "true"
// If set to non-empty string it must be comma-separated list of values that are all a subset
// of "EnableAllProposals" (takes precedence over ProposalsEnabled)
// https://github.com/CosmWasm/wasmd/blob/02a54d33ff2c064f3539ae12d75d027d9c665f05/x/wasm/internal/types/proposal.go#L28-L34
// EnableSpecificProposals = ""
)

// These constants are derived from the above variables.
// These are the ones we will want to use in the code, based on
// any overrides above
var (
	Bech32Prefix = AccountAddressPrefix

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32Prefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32Prefix + sdk.PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

var (
	Upgrades = []upgrades.Upgrade{v04.Upgrade, v05.Upgrade}
)

// GetEnabledProposals parses the ProposalsEnabled / EnableSpecificProposals values to
// produce a list of enabled proposals to pass into wasmd app.
/*func GetEnabledProposals() []wasmtypes.ProposalType {
	if EnableSpecificProposals == "" {
		if ProposalsEnabled == "true" {
			return wasmtypes.EnableAllProposals
		}
		return wasmtypes.DisableAllProposals
	}
	chunks := strings.Split(EnableSpecificProposals, ",")
	proposals, err := wasmtypes.ConvertToProposals(chunks)
	if err != nil {
		panic(err)
	}
	return proposals
}*/

func getGovProposalHandlers() []govclient.ProposalHandler {
	return []govclient.ProposalHandler{
		paramsclient.ProposalHandler,
		//upgradeclient.LegacyProposalHandler,
		//upgradeclient.LegacyCancelProposalHandler,
		//ibcclientclient.UpdateClientProposalHandler,
		//ibcclientclient.UpgradeProposalHandler,
	}
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// module account permissions
	maccPerms = map[string][]string{
		// SDK
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		govtypes.ModuleName:            {authtypes.Burner},
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		// IBC
		ibctransfertypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		icatypes.ModuleName:         nil,
		icqtypes.ModuleName:         nil,
		ibcfeetypes.ModuleName:      nil,
		// Wasm
		wasmtypes.ModuleName: {authtypes.Burner},
		// Custom
		tokenmoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner},
	}
)

var (
	_ servertypes.Application = (*App)(nil)
	_ runtime.AppI            = (*App)(nil)
)

func init() {
	/*var err error
	DefaultNodeHome, err = clienthelpers.GetNodeHomeDirectory(".app")*/
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry
	configurator      module.Configurator
	txConfig          client.TxConfig

	AppKeepers keepers.AppKeepers

	// Module Manager
	mm                 *module.Manager
	BasicModuleManager module.BasicManager

	// Simulation Manager
	sm *module.SimulationManager
}

func (app *App) setupUpgradeHandlers(cfg module.Configurator) {
	for _, upgrade := range Upgrades {
		app.AppKeepers.UpgradeKeeper.SetUpgradeHandler(
			upgrade.UpgradeName,
			upgrade.CreateUpgradeHandler(
				app.mm,
				cfg,
				&app.AppKeepers,
			),
		)
	}
}

// configure store loader that checks if version == upgradeHeight and applies store upgrades
func (app *App) setupUpgradeStoreLoaders() {
	upgradeInfo, err := app.AppKeepers.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic("failed to read upgrade info from disk" + err.Error())
	}

	if app.AppKeepers.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	for _, upgrade := range Upgrades {
		if upgradeInfo.Name == upgrade.UpgradeName {
			storeUpgrades := upgrade.StoreUpgrades
			app.SetStoreLoader(
				upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades),
			)
		}
	}
}

// New returns a reference to an initialized Firmachain.
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	/* skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig encparams.EncodingConfig,
	enabledProposals []wasmtypes.ProposalType,*/
	appOpts servertypes.AppOptions,
	wasmOpts []wasmkeeper.Option,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {

	legacyAmino := codec.NewLegacyAmino()
	interfaceRegistry, err := types.NewInterfaceRegistryWithOptions(types.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec: address.Bech32Codec{
				Bech32Prefix: Bech32PrefixAccAddr,
			},
			ValidatorAddressCodec: address.Bech32Codec{
				Bech32Prefix: Bech32PrefixValAddr,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	appCodec := codec.NewProtoCodec(interfaceRegistry)
	txConfig := authtx.NewTxConfig(appCodec, authtx.DefaultSignModes)

	std.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterInterfaces(interfaceRegistry)

	bApp := baseapp.NewBaseApp(Name, logger, db, txConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := storetypes.NewKVStoreKeys(
		// SDK
		authtypes.StoreKey,
		authzkeeper.StoreKey,
		banktypes.StoreKey,
		//capabilitytypes.StoreKey,
		consensusparamtypes.StoreKey,
		crisistypes.StoreKey,
		distrtypes.StoreKey,
		//evidencetypes.StoreKey,
		feegrant.StoreKey,
		govtypes.StoreKey,
		minttypes.StoreKey,
		paramstypes.StoreKey,
		slashingtypes.StoreKey,
		stakingtypes.StoreKey,
		upgradetypes.StoreKey,
		// IBC
		ibcexported.StoreKey,
		ibctransfertypes.StoreKey,
		ibcfeetypes.StoreKey,
		ibchookstypes.StoreKey,
		packetforwardtypes.StoreKey,
		icahosttypes.StoreKey,
		icacontrollertypes.StoreKey,
		icqtypes.StoreKey,
		// Wasm
		wasmtypes.StoreKey,
		// Custom
		contractmoduletypes.StoreKey,
		nftmoduletypes.StoreKey,
		tokenmoduletypes.StoreKey,
	)
	tkeys := storetypes.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		AppKeepers:        keepers.NewAppKeepersWithKeys(keys, tkeys, memKeys),
	}
	invCheckPeriod := cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod))
	skipUpgradeHeights := map[int64]bool{}
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}
	homePath := cast.ToString(appOpts.Get(flags.FlagHome))

	govModAddress := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	// ============ Keepers ============
	// TODO: Move keepers in app/keepers/keepers.go
	app.AppKeepers.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)
	scopedIBCKeeper := app.AppKeepers.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedTransferKeeper := app.AppKeepers.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedICAHostKeeper := app.AppKeepers.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	scopedICAControllerKeeper := app.AppKeepers.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	scopedICQKeeper := app.AppKeepers.CapabilityKeeper.ScopeToModule(icqtypes.ModuleName)
	scopedWasmKeeper := app.AppKeepers.CapabilityKeeper.ScopeToModule(wasmtypes.ModuleName)
	app.AppKeepers.ScopedIBCKeeper = scopedIBCKeeper
	app.AppKeepers.ScopedTransferKeeper = scopedTransferKeeper
	app.AppKeepers.ScopedICAHostKeeper = scopedICAHostKeeper
	app.AppKeepers.ScopedICAControllerKeeper = scopedICAControllerKeeper
	app.AppKeepers.ScopedICQKeeper = scopedICQKeeper
	app.AppKeepers.ScopedWasmKeeper = scopedWasmKeeper
	app.AppKeepers.ParamsKeeper = initParamsKeeper(
		appCodec,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)
	app.AppKeepers.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[consensusparamtypes.StoreKey]),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		runtime.EventService{},
	)
	bApp.SetParamStore(app.AppKeepers.ConsensusParamsKeeper.ParamsStore)
	app.AppKeepers.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		Bech32Prefix,
		govModAddress,
	)
	app.AppKeepers.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[banktypes.StoreKey]),
		app.AppKeepers.AccountKeeper,
		app.ModuleAccountAddrsForBankModule(),
		govModAddress,
		logger,
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[stakingtypes.StoreKey]),
		app.AppKeepers.AccountKeeper,
		app.AppKeepers.BankKeeper,
		govModAddress,
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	)
	app.AppKeepers.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[minttypes.StoreKey]),
		stakingKeeper,
		app.AppKeepers.AccountKeeper,
		app.AppKeepers.BankKeeper,
		authtypes.FeeCollectorName,
		govModAddress,
	)
	app.AppKeepers.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[distrtypes.StoreKey]),
		app.AppKeepers.AccountKeeper,
		app.AppKeepers.BankKeeper,
		stakingKeeper,
		authtypes.FeeCollectorName,
		govModAddress,
	)
	app.AppKeepers.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		legacyAmino,
		runtime.NewKVStoreService(keys[slashingtypes.StoreKey]),
		stakingKeeper,
		govModAddress,
	)
	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			app.AppKeepers.DistrKeeper.Hooks(),
			app.AppKeepers.SlashingKeeper.Hooks(),
		),
	)
	app.AppKeepers.StakingKeeper = stakingKeeper
	govKeeper := govkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[govtypes.StoreKey]),
		app.AppKeepers.AccountKeeper,
		app.AppKeepers.BankKeeper,
		app.AppKeepers.StakingKeeper,
		app.AppKeepers.DistrKeeper,
		bApp.MsgServiceRouter(),
		govtypes.DefaultConfig(),
		govModAddress,
	)
	app.AppKeepers.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register governance hooks
		),
	)
	app.AppKeepers.EvidenceKeeper = *evidencekeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[evidencetypes.StoreKey]),
		app.AppKeepers.StakingKeeper,
		app.AppKeepers.SlashingKeeper,
		authcodec.NewBech32Codec(sdk.Bech32PrefixAccAddr),
		runtime.ProvideCometInfoService(),
	)
	app.AppKeepers.CrisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[crisistypes.StoreKey]),
		invCheckPeriod,
		app.AppKeepers.BankKeeper,
		authtypes.FeeCollectorName,
		govModAddress,
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
	)
	app.AppKeepers.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[feegrant.StoreKey]),
		app.AppKeepers.AccountKeeper,
	)
	app.AppKeepers.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		runtime.NewKVStoreService(keys[upgradetypes.StoreKey]),
		appCodec,
		homePath,
		app.BaseApp,
		govModAddress,
	)
	app.AppKeepers.AuthzKeeper = authzkeeper.NewKeeper(
		runtime.NewKVStoreService(keys[authzkeeper.StoreKey]),
		appCodec,
		app.BaseApp.MsgServiceRouter(),
		app.AppKeepers.AccountKeeper,
	)
	// ------------ IBC ------------
	app.AppKeepers.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibcexported.StoreKey],
		app.GetSubspace(ibcexported.ModuleName),
		app.AppKeepers.StakingKeeper,
		app.AppKeepers.UpgradeKeeper,
		scopedIBCKeeper,
		govModAddress,
	)
	// Initialize packet forward middleware router
	app.AppKeepers.PacketForwardKeeper = packetforwardkeeper.NewKeeper(
		appCodec,
		app.AppKeepers.GetKeys()[packetforwardtypes.StoreKey],
		//app.GetSubspace(packetforwardtypes.ModuleName),
		app.AppKeepers.TransferKeeper, // Will be zero-value here. Reference is set later on with SetTransferKeeper.
		app.AppKeepers.IBCKeeper.ChannelKeeper,
		app.AppKeepers.DistrKeeper,
		app.AppKeepers.BankKeeper,
		app.AppKeepers.IBCKeeper.ChannelKeeper,
		govModAddress,
	)
	app.AppKeepers.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		// The ICS4Wrapper is replaced by the PacketForwardKeeper instead of the channel so that sending can be overridden by the middleware
		app.AppKeepers.PacketForwardKeeper,
		app.AppKeepers.IBCKeeper.ChannelKeeper,
		app.AppKeepers.IBCKeeper.PortKeeper,
		app.AppKeepers.AccountKeeper,
		app.AppKeepers.BankKeeper,
		scopedTransferKeeper,
		govModAddress,
	)
	app.AppKeepers.PacketForwardKeeper.SetTransferKeeper(app.AppKeepers.TransferKeeper)
	hooksKeeper := ibchookskeeper.NewKeeper(
		app.AppKeepers.GetKeys()[ibchookstypes.StoreKey],
	)
	app.AppKeepers.IBCHooksKeeper = &hooksKeeper
	wasmHooks := ibc_hooks.NewWasmHooks(app.AppKeepers.IBCHooksKeeper, &app.AppKeepers.WasmKeeper, Bech32Prefix) // The contract keeper needs to be set later
	app.AppKeepers.Ics20WasmHooks = &wasmHooks
	app.AppKeepers.HooksICS4Wrapper = ibc_hooks.NewICS4Middleware(
		app.AppKeepers.IBCKeeper.ChannelKeeper,
		app.AppKeepers.Ics20WasmHooks,
	)
	app.AppKeepers.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec,
		keys[icahosttypes.StoreKey],
		app.GetSubspace(icahosttypes.SubModuleName),
		app.AppKeepers.HooksICS4Wrapper,
		app.AppKeepers.IBCKeeper.ChannelKeeper,
		app.AppKeepers.IBCKeeper.PortKeeper,
		app.AppKeepers.AccountKeeper,
		scopedICAHostKeeper,
		app.MsgServiceRouter(),
		govModAddress,
	)
	app.AppKeepers.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		appCodec, app.AppKeepers.GetKeys()[icacontrollertypes.StoreKey],
		app.GetSubspace(icacontrollertypes.SubModuleName),
		app.AppKeepers.IBCFeeKeeper, // use ics29 fee as ics4Wrapper in middleware stack
		app.AppKeepers.IBCKeeper.ChannelKeeper,
		app.AppKeepers.IBCKeeper.PortKeeper,
		scopedICAControllerKeeper,
		bApp.MsgServiceRouter(),
		govModAddress,
	)
	app.AppKeepers.ICQKeeper = icqkeeper.NewKeeper(
		appCodec,
		app.AppKeepers.GetKeys()[icqtypes.StoreKey],
		//app.GetSubspace(icqtypes.ModuleName),
		app.AppKeepers.IBCKeeper.ChannelKeeper, // may be replaced with middleware
		app.AppKeepers.IBCKeeper.ChannelKeeper,
		app.AppKeepers.IBCKeeper.PortKeeper,
		scopedICQKeeper,
		bApp.GRPCQueryRouter(),
		govModAddress,
	)
	// Do not use this middleware for anything except x/wasm requirement.
	// The spec currently requires new channels to be created, to use it.
	// We need to wait for Channel Upgradability before we can use this for any other middleware.
	app.AppKeepers.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		appCodec,
		app.AppKeepers.GetKeys()[ibcfeetypes.StoreKey],
		app.AppKeepers.HooksICS4Wrapper, // replaced with IBC middleware
		app.AppKeepers.IBCKeeper.ChannelKeeper,
		app.AppKeepers.IBCKeeper.PortKeeper,
		app.AppKeepers.AccountKeeper,
		app.AppKeepers.BankKeeper,
	)
	// ----------- Wasm -----------
	wasmDir := filepath.Join(homePath, "data")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic("error : read wasm config: " + err.Error())
	}
	supportedFeatures := []string{"iterator,staking,stargate"}
	app.AppKeepers.WasmKeeper = wasmkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[wasmtypes.StoreKey]),
		app.AppKeepers.AccountKeeper,
		app.AppKeepers.BankKeeper,
		app.AppKeepers.StakingKeeper,
		distrkeeper.NewQuerier(app.AppKeepers.DistrKeeper),
		app.AppKeepers.IBCFeeKeeper,
		app.AppKeepers.IBCKeeper.ChannelKeeper,
		app.AppKeepers.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		app.AppKeepers.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		supportedFeatures,
		govModAddress,
		wasmOpts...,
	)
	// ------------ Custom ------------
	app.AppKeepers.ContractKeeper = *contractmodulekeeper.NewKeeper(
		appCodec,
		keys[contractmoduletypes.StoreKey],
		keys[contractmoduletypes.MemStoreKey],
	)
	app.AppKeepers.NftKeeper = *nftmodulekeeper.NewKeeper(
		appCodec,
		keys[nftmoduletypes.StoreKey],
		keys[nftmoduletypes.MemStoreKey],
	)
	app.AppKeepers.TokenKeeper = *tokenmodulekeeper.NewKeeper(
		appCodec,
		keys[tokenmoduletypes.StoreKey],
		keys[tokenmoduletypes.MemStoreKey],
		app.AppKeepers.BankKeeper,
	)

	// ============ TxConfig ============
	txConfigOpts := authtx.ConfigOptions{
		EnabledSignModes:           append(authtx.DefaultSignModes, sigtypes.SignMode_SIGN_MODE_TEXTUAL),
		TextualCoinMetadataQueryFn: authtxconfig.NewBankKeeperCoinMetadataQueryFn(app.AppKeepers.BankKeeper),
	}
	txConfig, err = authtx.NewTxConfigWithOptions(
		appCodec,
		txConfigOpts,
	)
	if err != nil {
		panic(err)
	}
	app.txConfig = txConfig

	// NOTE: upgrade module is required to be prioritized
	app.mm.SetOrderPreBlockers(
		upgradetypes.ModuleName,
	)

	// ============ Routers ============
	// ------------ Gov ------------
	govRouter := govv1beta.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govv1beta.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.AppKeepers.ParamsKeeper))
	//AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.AppKeepers.UpgradeKeeper)).
	//AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.AppKeepers.IBCKeeper.ClientKeeper))
	// DEPRECATED: DO NOT USE
	//
	//if len(enabledProposals) != 0 {
	//	govRouter.AddRoute(wasmtypes.RouterKey, wasm.NewWasmProposalHandler(app.AppKeepers.WasmKeeper, enabledProposals))
	//}
	// Set legacy router for backwards compatibility with gov v1beta1
	app.AppKeepers.GovKeeper.SetLegacyRouter(govRouter)

	// ------------ Ibc ------------
	var transferStack porttypes.IBCModule
	transferStack = transfer.NewIBCModule(app.AppKeepers.TransferKeeper)
	transferStack = ibcfee.NewIBCMiddleware(transferStack, app.AppKeepers.IBCFeeKeeper)
	transferStack = ibc_hooks.NewIBCMiddleware(transferStack, &app.AppKeepers.HooksICS4Wrapper)
	transferStack = packetforward.NewIBCMiddleware(
		transferStack,
		app.AppKeepers.PacketForwardKeeper,
		0,
		packetforwardkeeper.DefaultForwardTransferPacketTimeoutTimestamp,
		packetforwardkeeper.DefaultRefundTransferPacketTimeoutTimestamp,
	)
	// initialize ICA module with mock module as the authentication module on the controller side
	var icaControllerStack porttypes.IBCModule
	icaControllerStack = icacontroller.NewIBCMiddleware(icaControllerStack, app.AppKeepers.ICAControllerKeeper)
	icaControllerStack = ibcfee.NewIBCMiddleware(icaControllerStack, app.AppKeepers.IBCFeeKeeper)
	// RecvPacket, message that originates from core IBC and goes down to app, the flow is:
	// channel.RecvPacket -> fee.OnRecvPacket -> icaHost.OnRecvPacket
	var icaHostStack porttypes.IBCModule
	icaHostStack = icahost.NewIBCModule(app.AppKeepers.ICAHostKeeper)
	icaHostStack = ibcfee.NewIBCMiddleware(icaHostStack, app.AppKeepers.IBCFeeKeeper)
	// Create fee enabled wasm ibc Stack
	var wasmStack porttypes.IBCModule
	wasmStack = wasm.NewIBCHandler(app.AppKeepers.WasmKeeper, app.AppKeepers.IBCKeeper.ChannelKeeper, app.AppKeepers.IBCFeeKeeper)
	wasmStack = ibcfee.NewIBCMiddleware(wasmStack, app.AppKeepers.IBCFeeKeeper)
	icqModule := icq.NewIBCModule(app.AppKeepers.ICQKeeper)
	ibcRouter := ibcporttypes.NewRouter().
		AddRoute(ibctransfertypes.ModuleName, transferStack).
		AddRoute(wasmtypes.ModuleName, wasmStack).
		AddRoute(icahosttypes.SubModuleName, icaHostStack).
		AddRoute(icacontrollertypes.SubModuleName, icaControllerStack).
		AddRoute(icqtypes.ModuleName, icqModule)
	app.AppKeepers.IBCKeeper.SetRouter(ibcRouter)

	// ============ Modules ============
	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	appModules := []module.AppModule{
		// SDK modules
		auth.NewAppModule(appCodec, app.AppKeepers.AccountKeeper, nil, app.GetSubspace(authtypes.ModuleName)),
		authzmodule.NewAppModule(appCodec, app.AppKeepers.AuthzKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.interfaceRegistry),
		bank.NewAppModule(appCodec, app.AppKeepers.BankKeeper, app.AppKeepers.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.AppKeepers.CapabilityKeeper, false),
		consensusparams.NewAppModule(appCodec, app.AppKeepers.ConsensusParamsKeeper),
		crisis.NewAppModule(app.AppKeepers.CrisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),
		distr.NewAppModule(appCodec, app.AppKeepers.DistrKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		evidence.NewAppModule(app.AppKeepers.EvidenceKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.FeeGrantKeeper, app.interfaceRegistry),
		genutil.NewAppModule(app.AppKeepers.AccountKeeper, app.AppKeepers.StakingKeeper, app.BaseApp, app.txConfig), //TODO: check again if app is correct
		vesting.NewAppModule(app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper),
		gov.NewAppModule(appCodec, &app.AppKeepers.GovKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.AppKeepers.MintKeeper, app.AppKeepers.AccountKeeper, nil, app.GetSubspace(minttypes.ModuleName)),
		params.NewAppModule(app.AppKeepers.ParamsKeeper),
		slashing.NewAppModule(appCodec, app.AppKeepers.SlashingKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.StakingKeeper, app.GetSubspace(slashingtypes.ModuleName), app.interfaceRegistry),
		staking.NewAppModule(appCodec, app.AppKeepers.StakingKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		upgrade.NewAppModule(app.AppKeepers.UpgradeKeeper, app.AppKeepers.AccountKeeper.AddressCodec()),
		vesting.NewAppModule(app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper),
		// IBC modules
		ibc.NewAppModule(app.AppKeepers.IBCKeeper),
		transfer.NewAppModule(app.AppKeepers.TransferKeeper),
		ibcfee.NewAppModule(app.AppKeepers.IBCFeeKeeper),
		ibc_hooks.NewAppModule(app.AppKeepers.AccountKeeper),
		packetforward.NewAppModule(app.AppKeepers.PacketForwardKeeper, app.GetSubspace(packetforwardtypes.ModuleName)),
		ica.NewAppModule(&app.AppKeepers.ICAControllerKeeper, &app.AppKeepers.ICAHostKeeper),
		icq.NewAppModule(app.AppKeepers.ICQKeeper, app.GetSubspace(icqtypes.ModuleName)),
		// Wasm modules
		wasm.NewAppModule(appCodec, &app.AppKeepers.WasmKeeper, app.AppKeepers.StakingKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.MsgServiceRouter(), app.GetSubspace(wasmtypes.ModuleName)),
		// Custom modules
		contractmodule.NewAppModule(appCodec, app.AppKeepers.ContractKeeper),
		nftmodule.NewAppModule(appCodec, app.AppKeepers.NftKeeper),
		tokenmodule.NewAppModule(appCodec, app.AppKeepers.TokenKeeper),
	}

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	orderBeginBlockers := []string{
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		vestingtypes.ModuleName,
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		authz.ModuleName,
		consensusparamtypes.ModuleName,
		nftmoduletypes.ModuleName,
		contractmoduletypes.ModuleName,
		tokenmoduletypes.ModuleName,
		wasmtypes.ModuleName,
		icatypes.ModuleName,
		packetforwardtypes.ModuleName,
		ibcfeetypes.ModuleName,
		icqtypes.ModuleName,
		ibchookstypes.ModuleName,
	}

	orderEndBlockers := []string{
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		vestingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		consensusparamtypes.ModuleName,
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		authz.ModuleName,
		nftmoduletypes.ModuleName,
		contractmoduletypes.ModuleName,
		tokenmoduletypes.ModuleName,
		wasmtypes.ModuleName,
		icatypes.ModuleName,
		packetforwardtypes.ModuleName,
		ibcfeetypes.ModuleName,
		icqtypes.ModuleName,
		ibchookstypes.ModuleName,
	}

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	orderInitGenesis := []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		vestingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibcexported.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		consensusparamtypes.ModuleName,
		ibctransfertypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		nftmoduletypes.ModuleName,
		contractmoduletypes.ModuleName,
		tokenmoduletypes.ModuleName,
		wasmtypes.ModuleName,
		icatypes.ModuleName,
		packetforwardtypes.ModuleName,
		ibcfeetypes.ModuleName,
		ibchookstypes.ModuleName,
		icqtypes.ModuleName,
	}

	app.mm = module.NewManager(appModules...)
	app.BasicModuleManager = newBasicManagerFromManager(app)
	app.mm.SetOrderBeginBlockers(orderBeginBlockers...)
	app.mm.SetOrderEndBlockers(orderEndBlockers...)
	app.mm.SetOrderInitGenesis(orderInitGenesis...)

	app.mm.RegisterInvariants(app.AppKeepers.CrisisKeeper)

	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// create the simulation manager and define the order of the modules for deterministic simulations
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AppKeepers.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		bank.NewAppModule(appCodec, app.AppKeepers.BankKeeper, app.AppKeepers.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.AppKeepers.CapabilityKeeper, false),
		feegrantmodule.NewAppModule(appCodec, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AppKeepers.AuthzKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, &app.AppKeepers.GovKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.AppKeepers.MintKeeper, app.AppKeepers.AccountKeeper, nil, app.GetSubspace(minttypes.ModuleName)),
		staking.NewAppModule(appCodec, app.AppKeepers.StakingKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		distr.NewAppModule(appCodec, app.AppKeepers.DistrKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.AppKeepers.SlashingKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.StakingKeeper, app.GetSubspace(stakingtypes.ModuleName), app.interfaceRegistry),
		params.NewAppModule(app.AppKeepers.ParamsKeeper),
		evidence.NewAppModule(app.AppKeepers.EvidenceKeeper),
		wasm.NewAppModule(appCodec, &app.AppKeepers.WasmKeeper, app.AppKeepers.StakingKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.MsgServiceRouter(), app.GetSubspace(wasmtypes.ModuleName)),
		ibc.NewAppModule(app.AppKeepers.IBCKeeper),
		transfer.NewAppModule(app.AppKeepers.TransferKeeper),
		ibcfee.NewAppModule(app.AppKeepers.IBCFeeKeeper),
		// actually we don't use simulation manger yet.
		// contractmodule.NewAppModule(appCodec, app.AppKeepers.ContractKeeper),
		// nftmodule.NewAppModule(appCodec, app.AppKeepers.NftKeeper),
		// tokenmodule.NewAppModule(appCodec, app.AppKeepers.TokenKeeper),
	)
	app.sm.RegisterStoreDecoders()

	// ============ Initialization ============
	app.setupUpgradeHandlers(app.configurator)

	// SDK v47 - since we do not use dep inject, this gives us access to newer gRPC services.
	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.mm.Modules))
	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetPreBlocker(app.PreBlocker)
	app.SetBeginBlocker(app.BeginBlocker)

	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.AppKeepers.AccountKeeper,
			BankKeeper:      app.AppKeepers.BankKeeper,
			SignModeHandler: app.txConfig.SignModeHandler(),
			FeegrantKeeper:  app.AppKeepers.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)

	// In v0.46, the SDK introduces _postHandlers_. PostHandlers are like
	// antehandlers, but are run _after_ the `runMsgs` execution. They are also
	// defined as a chain, and have the same signature as antehandlers.
	//
	// In baseapp, postHandlers are run in the same store branch as `runMsgs`,
	// meaning that both `runMsgs` and `postHandler` state will be committed if
	// both are successful, and both will be reverted if any of the two fails.
	//
	// The SDK exposes a default postHandlers chain, which comprises of only
	// one decorator: the Transaction Tips decorator. However, some chains do
	// not need it by default, so feel free to comment the next line if you do
	// not need tips.
	// To read more about tips:
	// https://docs.cosmos.network/main/core/tips.html
	//
	// Please note that changing any of the anteHandler or postHandler chain is
	// likely to be a state-machine breaking change, which needs a coordinated
	// upgrade.
	app.setPostHandler()

	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			ctmos.Exit(err.Error())
		}

		// Initialize and seal the capability keeper so all persistent capabilities
		// are loaded in-memory and prevent any further modules from creating scoped
		// sub-keepers.
		// This must be done during creation of baseapp rather than in InitChain so
		// that in-memory capabilities get regenerated on app restart.
		// Note that since this reads from the store, we can only perform it when
		// `loadLatest` is set to true.
		app.AppKeepers.CapabilityKeeper.Seal()
	}

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// GetBaseApp returns the base app of the application
func (app App) GetBaseApp() *baseapp.BaseApp { return app.BaseApp }

// PreBlocker application updates every pre block
func (app *App) PreBlocker(ctx sdk.Context, _ *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	return app.mm.PreBlock(ctx)
}

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	app.updateValidatorMinCommision(ctx)
	return app.mm.BeginBlock(ctx)
}

func (app *App) updateValidatorMinCommision(ctx sdk.Context) {
	staking := app.AppKeepers.StakingKeeper

	validators, err := staking.GetAllValidators(ctx)
	if err != nil {
		panic(err) //TODO: handle
	}
	minCommissionRate := math.LegacyMustNewDecFromStr("0.05")

	for _, v := range validators {
		//
		if v.Commission.Rate.LT(minCommissionRate) {
			v.Commission.Rate = minCommissionRate
			v.Commission.UpdateTime = ctx.BlockHeader().Time

			valBs, err := staking.ValidatorAddressCodec().StringToBytes(v.GetOperator())
			if err != nil {
				panic(err)
			}
			// call the before-modification hook since we're about to update the commission
			staking.Hooks().BeforeValidatorModified(ctx, valBs)
			staking.SetValidator(ctx, v)
		}
	}
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	return app.mm.EndBlock(ctx)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	var genesisState GenesisState
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.AppKeepers.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

func (app *App) ModuleAccountAddrsForBankModule() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}
	return modAccAddrs
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// ModuleBasics defines the module BasicManager that is in charge of setting up basic,
// non-dependant module elements, such as codec registration
// and genesis verification.
func newBasicManagerFromManager(app *App) module.BasicManager {
	basicManager := module.NewBasicManagerFromManager(
		app.mm,
		map[string]module.AppModuleBasic{
			genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
			govtypes.ModuleName: gov.NewAppModuleBasic(
				[]govclient.ProposalHandler{
					paramsclient.ProposalHandler,
				},
			),
		})
	basicManager.RegisterLegacyAminoCodec(app.legacyAmino)
	basicManager.RegisterInterfaces(app.interfaceRegistry)
	return basicManager
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Gaia's InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.AppKeepers.GetKeys()[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.AppKeepers.GetTkeys()[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.AppKeepers.GetMemKeys()[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.AppKeepers.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx

	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// TODO: enable?
	// Register new tendermint queries routes from grpc-gateway.
	// tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	app.BasicModuleManager.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register app's OpenAPI routes.
	apiSvr.Router.Handle("/static/openapi.yml", http.FileServer(http.FS(docs.Docs)))
	apiSvr.Router.HandleFunc("/", openapiconsole.Handler(Name, "/static/openapi.yml"))
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	cmtservice.RegisterTendermintService(clientCtx, app.BaseApp.GRPCQueryRouter(), app.interfaceRegistry, app.Query)
}

func (app *App) RegisterNodeService(clientCtx client.Context, cfg config.Config) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter(), cfg)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)
	// SDK
	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	// IBC
	paramsKeeper.Subspace(ibcexported.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(packetforwardtypes.ModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	// Wasm
	paramsKeeper.Subspace(wasmtypes.ModuleName)
	// Custom
	paramsKeeper.Subspace(contractmoduletypes.ModuleName)
	paramsKeeper.Subspace(nftmoduletypes.ModuleName)
	paramsKeeper.Subspace(tokenmoduletypes.ModuleName)

	return paramsKeeper
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

func (app *App) setPostHandler() {
	postHandler, err := posthandler.NewPostHandler(
		posthandler.HandlerOptions{},
	)
	if err != nil {
		panic(err)
	}

	app.SetPostHandler(postHandler)
}
