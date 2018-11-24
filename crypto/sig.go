package crypto

import (
	"encoding/hex"
	"errors"
)

var (
	// ErrInvalidHexSig is returned if the hexadecimal representation couldn't be decoded
	ErrInvalidHexSig = errors.New("invalid hex representation of signature")
	// ErrInvalidSigLen is returned if the signature length is different than ed25519 64 byte
	ErrInvalidSigLen = errors.New("invalid signature length , signatures are exactly 64 bytes")
)

// Signature type
type Signature [64]byte

// NewSignature returns a new signature using a byte container
func NewSignature(b []byte) (Signature, error) {
	sig := Signature{}

	if len(b) != len(sig) {
		return Signature{}, ErrInvalidSigLen
	}

	copy(sig[:], b[:])

	return sig, nil

}

// SigFromHex converts a hex encoded string that represent a ed signature to a signature type
func SigFromHex(s string) (Signature, error) {
	b, err := hex.DecodeString(s)

	if err != nil {
		return Signature{}, ErrInvalidHexSig
	}

	return NewSignature(b)

}
