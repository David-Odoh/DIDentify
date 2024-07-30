package keeper

import (
	"context"
	"fmt"

	"identity/x/authentication/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AuthenticateUser(goCtx context.Context, msg *types.MsgAuthenticateUser) (*types.MsgAuthenticateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valid, err := k.VerifyChallenge(ctx, msg.Did, msg.SignedChallenge)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, fmt.Errorf("authentication failed")
	}

	return &types.MsgAuthenticateUserResponse{
		Authenticated: true,
	}, nil
}
