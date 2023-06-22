package keeper

import (
	"github.com/sonrhq/sonr/x/service/types"
)

var _ types.QueryServer = Keeper{}
