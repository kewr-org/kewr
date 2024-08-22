package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kewr/testutil/keeper"
	"kewr/x/wms/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.WmsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
