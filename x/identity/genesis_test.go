package identity_test

import (
	"testing"

	keepertest "github.com/sonrhq/sonr/testutil/keeper"
	"github.com/sonrhq/sonr/testutil/nullify"
	"github.com/sonrhq/sonr/x/identity"
	"github.com/sonrhq/sonr/x/identity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		DIDDocumentList: []types.DIDDocument{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, *k, genesisState)
	got := identity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.DIDDocumentList, got.DIDDocumentList)
	// this line is used by starport scaffolding # genesis/test/assert
}
