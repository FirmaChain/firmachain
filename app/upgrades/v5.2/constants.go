package v5_2

import (
	store "cosmossdk.io/store/types"
	upgrades "github.com/firmachain/firmachain/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the upgrade.
const UpgradeName = "v0.5.2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV0_5_2UpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}
