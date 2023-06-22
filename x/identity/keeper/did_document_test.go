package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonrhq/sonr/testutil/keeper"
	"github.com/sonrhq/sonr/testutil/nullify"
	"github.com/sonrhq/sonr/x/identity/keeper"
	"github.com/sonrhq/sonr/x/identity/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDIDDocument(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DIDDocument {
	items := make([]types.DIDDocument, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetDIDDocument(ctx, items[i])
	}
	return items
}

func TestDIDDocumentGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDIDDocument(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDIDDocument(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDIDDocumentRemove(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDIDDocument(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDIDDocument(ctx,
			item.Index,
		)
		_, found := keeper.GetDIDDocument(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestDIDDocumentGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDIDDocument(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDIDDocument(ctx)),
	)
}
