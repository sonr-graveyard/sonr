package service_test

import (
	"testing"

	keepertest "github.com/sonrhq/sonr/testutil/keeper"
	"github.com/sonrhq/sonr/testutil/nullify"
	"github.com/sonrhq/sonr/x/service"
	"github.com/sonrhq/sonr/x/service/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ServiceKeeper(t)
	service.InitGenesis(ctx, *k, genesisState)
	got := service.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
