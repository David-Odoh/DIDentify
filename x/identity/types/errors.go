package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/identity module sentinel errors
var (
	ErrInvalidSigner    = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample           = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrIdentityNotFound = sdkerrors.Register(ModuleName, 1102, "Identity not found")
	ErrAlreadyVerified  = sdkerrors.Register(ModuleName, 1103, "Identity already verified")
	WrongHash           = sdkerrors.Register(ModuleName, 1104, "NationalID hashes not matching, wrong NationalID")
	DidExists           = sdkerrors.Register(ModuleName, 1105, "Did already associated with identity")
)
