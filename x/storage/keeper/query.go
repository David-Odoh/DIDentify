package keeper

import (
	"identity/x/storage/types"
)

var _ types.QueryServer = Keeper{}
