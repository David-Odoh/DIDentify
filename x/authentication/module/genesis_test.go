package authentication_test

import (
	"testing"

	keepertest "identity/testutil/keeper"
	"identity/testutil/nullify"
	authentication "identity/x/authentication/module"
	"identity/x/authentication/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AuthenticationKeeper(t)
	authentication.InitGenesis(ctx, k, genesisState)
	got := authentication.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
