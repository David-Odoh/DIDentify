package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterIdentity{}

func NewMsgRegisterIdentity(creator string, name string, birthdate string, nationalId string) *MsgRegisterIdentity {
	return &MsgRegisterIdentity{
		Creator:    creator,
		Name:       name,
		Birthdate:  birthdate,
		NationalId: nationalId,
	}
}

func (msg *MsgRegisterIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
