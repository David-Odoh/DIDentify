package testmod_test

import (
	"testing"

	keepertest "identity/testutil/keeper"
	"identity/testutil/nullify"
	testmod "identity/x/testmod/module"
	"identity/x/testmod/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestmodKeeper(t)
	testmod.InitGenesis(ctx, k, genesisState)
	got := testmod.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
