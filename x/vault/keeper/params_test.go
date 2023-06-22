package keeper_test

import (
	"testing"

	testkeeper "github.com/sonrhq/sonr/testutil/keeper"
	"github.com/sonrhq/sonr/x/vault/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VaultKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
