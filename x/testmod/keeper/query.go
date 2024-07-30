package keeper

import (
	"identity/x/testmod/types"
)

var _ types.QueryServer = Keeper{}
