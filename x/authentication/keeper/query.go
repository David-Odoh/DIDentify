package keeper

import (
	"identity/x/authentication/types"
)

var _ types.QueryServer = Keeper{}
