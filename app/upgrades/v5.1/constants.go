package v5_1

import (
	store "cosmossdk.io/store/types"
	upgrades "github.com/firmachain/firmachain/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the upgrade.
const UpgradeName = "v0.5.1"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV0_5_1UpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}
