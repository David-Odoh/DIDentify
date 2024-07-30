package keeper

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"identity/x/authentication/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetChallenge stores the challenge associated with a DID
func (k Keeper) SetChallenge(ctx context.Context, did string, challenge string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ChallengeKey))
	store.Set([]byte(did), []byte(challenge))
}

// GetChallenge retrieves the challenge associated with a DID
func (k Keeper) GetChallenge(ctx context.Context, did string) (string, bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ChallengeKey))
	bz := store.Get([]byte(did))
	if bz == nil {
		return "", false
	}
	return string(bz), true
}

// GenerateChallenge creates a new cryptographic challenge and stores it associated with the DID
func (k Keeper) GenerateChallenge(ctx context.Context, did string) (string, error) {
	challenge := make([]byte, 32)
	_, err := rand.Read(challenge)
	if err != nil {
		return "", err
	}
	challengeStr := hex.EncodeToString(challenge)
	k.SetChallenge(ctx, did, challengeStr)
	return challengeStr, nil
}

// VerifyChallenge verifies the signed challenge against the expected challenge
func (k Keeper) VerifyChallenge(ctx context.Context, did string, signedChallenge string) (bool, error) {
	storedChallenge, found := k.GetChallenge(ctx, did)
	if !found {
		return false, fmt.Errorf("challenge not found for DID")
	}

	publicKeyBytes, err := hex.DecodeString(did)
	if err != nil {
		return false, fmt.Errorf("invalid DID format: %v", err)
	}

	pubKey, err := decodeECDSAPublicKey(publicKeyBytes)
	if err != nil {
		return false, fmt.Errorf("failed to decode public key: %v", err)
	}

	signature, err := hex.DecodeString(signedChallenge)
	if err != nil {
		return false, fmt.Errorf("invalid signature format: %v", err)
	}

	hash := sha256.Sum256([]byte(storedChallenge))
	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])
	valid := ecdsa.Verify(pubKey, hash[:], r, s)

	return valid, nil
}

func decodeECDSAPublicKey(pubKeyBytes []byte) (*ecdsa.PublicKey, error) {
	if len(pubKeyBytes) != 65 {
		return nil, fmt.Errorf("invalid public key length")
	}
	pubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     new(big.Int).SetBytes(pubKeyBytes[1:33]),
		Y:     new(big.Int).SetBytes(pubKeyBytes[33:]),
	}
	return pubKey, nil
}
