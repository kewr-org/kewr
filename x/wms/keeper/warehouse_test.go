package keeper_test

import (
	"context"
	"testing"

	keepertest "kewr/testutil/keeper"
	"kewr/testutil/nullify"
	"kewr/x/wms/keeper"
	"kewr/x/wms/types"

	"github.com/stretchr/testify/require"
)

func createNWarehouse(keeper keeper.Keeper, ctx context.Context, n int) []types.Warehouse {
	items := make([]types.Warehouse, n)
	for i := range items {
		items[i].Id = keeper.AppendWarehouse(ctx, items[i])
	}
	return items
}

func TestWarehouseGet(t *testing.T) {
	keeper, ctx := keepertest.WmsKeeper(t)
	items := createNWarehouse(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetWarehouse(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestWarehouseRemove(t *testing.T) {
	keeper, ctx := keepertest.WmsKeeper(t)
	items := createNWarehouse(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWarehouse(ctx, item.Id)
		_, found := keeper.GetWarehouse(ctx, item.Id)
		require.False(t, found)
	}
}

func TestWarehouseGetAll(t *testing.T) {
	keeper, ctx := keepertest.WmsKeeper(t)
	items := createNWarehouse(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWarehouse(ctx)),
	)
}

func TestWarehouseCount(t *testing.T) {
	keeper, ctx := keepertest.WmsKeeper(t)
	items := createNWarehouse(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetWarehouseCount(ctx))
}
