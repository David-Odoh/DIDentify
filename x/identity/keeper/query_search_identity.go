package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"identity/x/identity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SearchIdentity(goCtx context.Context, req *types.QuerySearchIdentityRequest) (*types.QuerySearchIdentityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	hash := sha256.Sum256([]byte(req.NationalId))
	nationalIDhash := hex.EncodeToString(hash[:])

	identity, found := k.SearchByNationalID(ctx, nationalIDhash)
	if !found {
		return nil, status.Error(codes.NotFound, "identity not found")
	}

	return &types.QuerySearchIdentityResponse{Identity: &identity}, nil
}
