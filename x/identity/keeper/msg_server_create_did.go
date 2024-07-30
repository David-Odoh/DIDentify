package keeper

import (
	"context"
	"fmt"

	"identity/x/identity/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDid(goCtx context.Context, msg *types.MsgCreateDid) (*types.MsgCreateDidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Replace with the appropriate wallet UUID
	// WalletIdentifier := "b40a3b29-2a55-4db5-84ff-ff60380635eb"
	did, err := k.CreateDIDHandler(goCtx, msg.WalletIdentifier)
	if err != nil {
		return nil, err
	}

	identity, found := k.GetIdentity(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrIdentityNotFound, "identity with ID %d not found", msg.Id)
	}

	fmt.Println("Context:", ctx)
	fmt.Println("Creator:", msg.Creator)
	fmt.Println("DID:", did)

	if len(identity.Did) != 0 {
		return nil, errorsmod.Wrapf(types.DidExists, "Hi %s, you already have a DID", identity.Name)
	}

	identity.Did = did
	k.SetIdentity(ctx, identity)

	return &types.MsgCreateDidResponse{}, nil
}
