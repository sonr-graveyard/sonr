package keeper

import (
	"github.com/sonrhq/sonr/x/vault/types"
)

var _ types.QueryServer = Keeper{}
