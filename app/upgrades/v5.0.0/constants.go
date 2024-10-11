package v500

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/neutron-org/neutron/v5/app/upgrades"
	ibcratelimittypes "github.com/neutron-org/neutron/v5/x/ibc-rate-limit/types"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "v5.0.0"

	MarketMapAuthorityMultisig = "neutron1anjpluecd0tdc0n8xzc3l5hua4h93wyq0x7v56"
	// RateLimitContract defines the RL contract addr which we set as a contract address in ibc-rate-limit middleware
	// https://neutron.celat.one/neutron-1/contracts/neutron15aqgplxcavqhurr0g5wwtdw6025dknkqwkfh0n46gp2qjl6236cs2yd3nl
	RateLimitContract = "neutron15aqgplxcavqhurr0g5wwtdw6025dknkqwkfh0n46gp2qjl6236cs2yd3nl"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{ibcratelimittypes.ModuleName},
	},
}
