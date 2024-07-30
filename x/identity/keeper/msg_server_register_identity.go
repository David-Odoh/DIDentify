package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"identity/x/identity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterIdentity(goCtx context.Context, msg *types.MsgRegisterIdentity) (*types.MsgRegisterIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hash := sha256.Sum256([]byte(msg.NationalId))
	nationalIDhash := hex.EncodeToString(hash[:])

	// Create a new Identity object
	var identity = types.Identity{
		Creator:    msg.Creator,
		Name:       msg.Name,
		Birthdate:  msg.Birthdate,
		NationalId: nationalIDhash,
		Verified:   false,
		Did:        msg.Did,
	}

	// Store the identity
	k.AppendIdentity(ctx, identity)

	return &types.MsgRegisterIdentityResponse{}, nil
}
