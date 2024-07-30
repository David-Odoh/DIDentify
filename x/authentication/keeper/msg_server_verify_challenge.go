package keeper

import (
	"context"

	"identity/x/authentication/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VerifyChallenge(goCtx context.Context, msg *types.MsgVerifyChallenge) (*types.MsgVerifyChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valid, err := k.Keeper.VerifyChallenge(ctx, msg.Did, msg.SignedChallenge)
	if err != nil {
		return nil, err
	}

	return &types.MsgVerifyChallengeResponse{Valid: valid}, nil
}
