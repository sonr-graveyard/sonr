package keeper

import (
	"github.com/sonrhq/sonr/x/identity/types"
)

var _ types.QueryServer = Keeper{}
