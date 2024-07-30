package keeper

import (
	"context"

	"identity/x/authentication/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) GenerateChallenge(goCtx context.Context, msg *types.MsgGenerateChallenge) (*types.MsgGenerateChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	challenge, err := k.Keeper.GenerateChallenge(ctx, msg.Did)
	if err != nil {
		return nil, err
	}

	return &types.MsgGenerateChallengeResponse{Challenge: challenge}, nil
}
