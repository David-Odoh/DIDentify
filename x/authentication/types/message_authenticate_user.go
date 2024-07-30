package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAuthenticateUser{}

func NewMsgAuthenticateUser(creator string, did string, signedChallenge string) *MsgAuthenticateUser {
	return &MsgAuthenticateUser{
		Creator:         creator,
		Did:             did,
		SignedChallenge: signedChallenge,
	}
}

func (msg *MsgAuthenticateUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
