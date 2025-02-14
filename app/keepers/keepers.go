package keepers

import (
	storetypes "cosmossdk.io/store/types"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"

	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"

	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"

	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"

	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	contractmodulekeeper "github.com/firmachain/firmachain/v05/x/contract/keeper"
	nftmodulekeeper "github.com/firmachain/firmachain/v05/x/nft/keeper"

	tokenmodulekeeper "github.com/firmachain/firmachain/v05/x/token/keeper"

	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	packetforwardkeeper "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/keeper"
	icqkeeper "github.com/cosmos/ibc-apps/modules/async-icq/v8/keeper"
	ibc_hooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8"
	ibchookskeeper "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/keeper"
	ibcfeekeeper "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/keeper"

	circuitkeeper "cosmossdk.io/x/circuit/keeper"
)

type AppKeepers struct {
	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// SDK Keepers
	AccountKeeper         authkeeper.AccountKeeper
	AuthzKeeper           authzkeeper.Keeper
	BankKeeper            bankkeeper.Keeper
	CircuitKeeper         circuitkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	GovKeeper             govkeeper.Keeper
	MintKeeper            mintkeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper

	// IBC Keepers
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	TransferKeeper      ibctransferkeeper.Keeper
	IBCFeeKeeper        ibcfeekeeper.Keeper
	IBCHooksKeeper      *ibchookskeeper.Keeper
	PacketForwardKeeper *packetforwardkeeper.Keeper
	ICAHostKeeper       *icahostkeeper.Keeper
	ICAControllerKeeper *icacontrollerkeeper.Keeper
	ICQKeeper           icqkeeper.Keeper

	// Wasm Keepers
	WasmKeeper wasmkeeper.Keeper

	// Custom Keepers
	NftKeeper      nftmodulekeeper.Keeper
	ContractKeeper contractmodulekeeper.Keeper
	TokenKeeper    tokenmodulekeeper.Keeper

	// Scoped Keepers (public for test purposes)
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICQKeeper           capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper          capabilitykeeper.ScopedKeeper

	// IBC Hooks
	Ics20WasmHooks   *ibc_hooks.WasmHooks
	HooksICS4Wrapper ibc_hooks.ICS4Middleware
}

func NewAppKeepersWithKeys(
	keys map[string]*storetypes.KVStoreKey,
	tkeys map[string]*storetypes.TransientStoreKey,
	memKeys map[string]*storetypes.MemoryStoreKey,
) AppKeepers {
	return AppKeepers{
		keys:    keys,
		tkeys:   tkeys,
		memKeys: memKeys,
	}
}

// TODO: remove these once keepers are defined in the module
func (appKeepers *AppKeepers) GetKeys() map[string]*storetypes.KVStoreKey {
	return appKeepers.keys
}
func (appKeepers *AppKeepers) GetTkeys() map[string]*storetypes.TransientStoreKey {
	return appKeepers.tkeys
}
func (appKeepers *AppKeepers) GetMemKeys() map[string]*storetypes.MemoryStoreKey {
	return appKeepers.memKeys
}
