package v05

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7/router/types"
	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v7/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	upgrades "github.com/firmachain/firmachain/v05/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the upgrade.
const UpgradeName = "v0.5.0"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV0_5_0UpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			// new modules
			ibcfeetypes.StoreKey,
			ibchookstypes.StoreKey,
			packetforwardtypes.StoreKey,
			icahosttypes.StoreKey,
			icacontrollertypes.StoreKey,
		},
	},
}
