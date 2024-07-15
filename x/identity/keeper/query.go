package keeper

import (
	"identity/x/identity/types"
)

var _ types.QueryServer = Keeper{}
