package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		WarehouseList: []Warehouse{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in warehouse
	warehouseIdMap := make(map[uint64]bool)
	warehouseCount := gs.GetWarehouseCount()
	for _, elem := range gs.WarehouseList {
		if _, ok := warehouseIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for warehouse")
		}
		if elem.Id >= warehouseCount {
			return fmt.Errorf("warehouse id should be lower or equal than the last id")
		}
		warehouseIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
