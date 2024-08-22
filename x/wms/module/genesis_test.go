package wms_test

import (
	"testing"

	keepertest "kewr/testutil/keeper"
	"kewr/testutil/nullify"
	wms "kewr/x/wms/module"
	"kewr/x/wms/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WarehouseList: []types.Warehouse{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		WarehouseCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WmsKeeper(t)
	wms.InitGenesis(ctx, k, genesisState)
	got := wms.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.WarehouseList, got.WarehouseList)
	require.Equal(t, genesisState.WarehouseCount, got.WarehouseCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
