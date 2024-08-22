package wms

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"kewr/x/wms/keeper"
	"kewr/x/wms/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the warehouse
	for _, elem := range genState.WarehouseList {
		k.SetWarehouse(ctx, elem)
	}

	// Set warehouse count
	k.SetWarehouseCount(ctx, genState.WarehouseCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.WarehouseList = k.GetAllWarehouse(ctx)
	genesis.WarehouseCount = k.GetWarehouseCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
