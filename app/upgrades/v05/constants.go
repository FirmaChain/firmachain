package v05

import (
	store "cosmossdk.io/store/types"
	circuittypes "cosmossdk.io/x/circuit/types"
	consensusparamstypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistype "github.com/cosmos/cosmos-sdk/x/crisis/types"
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v8/types"
	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/types"

	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	upgrades "github.com/firmachain/firmachain/v05/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the upgrade.
const UpgradeName = "v0.5.0"

const legacyBurnModuleStoreKey = "burn"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV0_5_0UpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			// new modules
			ibcfeetypes.StoreKey,
			ibchookstypes.StoreKey,
			packetforwardtypes.StoreKey,
			icacontrollertypes.StoreKey,
			circuittypes.StoreKey,
			consensusparamstypes.StoreKey,
			//TODO: Why is crisis needed here?
			crisistype.StoreKey,
			icqtypes.StoreKey,
		},
		Deleted: []string{
			legacyBurnModuleStoreKey,
			// TODO: capability?
			// TODO: evidence?
		},
	},
}
