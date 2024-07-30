package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "identity/testutil/keeper"
	"identity/x/authentication/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AuthenticationKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
