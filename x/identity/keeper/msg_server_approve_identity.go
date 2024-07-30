package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"identity/x/identity/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApproveIdentity(goCtx context.Context, msg *types.MsgApproveIdentity) (*types.MsgApproveIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hash := sha256.Sum256([]byte(msg.NationalId))
	nationalIDhash := hex.EncodeToString(hash[:])

	identity, found := k.GetIdentity(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrIdentityNotFound, "key %d doesn't exist", msg.Id)
	}
	if identity.Verified != false {
		return nil, errorsmod.Wrapf(types.ErrAlreadyVerified, "Identity not pending", msg.Id)
	}

	if identity.NationalId != nationalIDhash {
		return nil, errorsmod.Wrapf(types.WrongHash, "hash %d is wrong", msg.Id)
	}

	identity.Verified = true
	k.SetIdentity(ctx, identity)
	return &types.MsgApproveIdentityResponse{}, nil
}
