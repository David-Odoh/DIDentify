package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "identity/testutil/keeper"
	"identity/x/testmod/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TestmodKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
