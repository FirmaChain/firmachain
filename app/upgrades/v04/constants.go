package v04

import (
	store "cosmossdk.io/store/types"
	upgrades "github.com/firmachain/firmachain/v05/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the upgrade.
const UpgradeName = "v0.4.0"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV0_4_0UpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}
