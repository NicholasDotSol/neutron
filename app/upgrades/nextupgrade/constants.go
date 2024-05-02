package nextupgrade

import (
	storetypes "cosmossdk.io/store/types"
	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"

	"github.com/neutron-org/neutron/v4/app/upgrades"
	globalfeetypes "github.com/neutron-org/neutron/v4/x/globalfee/types"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "nextupgrade"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			globalfeetypes.ModuleName,
			marketmaptypes.ModuleName,
			oracletypes.ModuleName,
		},
	},
}
