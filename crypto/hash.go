// Copyright 2018 The Cloq Authors
// This file is part of the cloq-core library .
//
// The cloq-core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cloq-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cloq-core library. If not, see <http://www.gnu.org/licenses/>.

// Hash provides useful functions to deal with Hashes

package crypto

import (
	"math/big"

	"golang.org/x/crypto/blake2b"
)

// HashLength as defined by the Blake2b Spec
const (
	HashLength      = 32             // 32 bytes is the length of Blake2b256 Hashes as used by cloq
	MaxHashLength   = HashLength * 2 // 64 bytes
	SmallHashLength = 20             // 20 bytes for 160 bit hashes
)

// Hash represents the 32 byte digest of Blake2b
type Hash [HashLength]byte

// NilHash returns the zero hash
func NilHash() Hash {

	var zerohash [32]byte
	return zerohash

}

// HashString returns a string as a 32byte array
func HashString(s string) Hash {

	h := NilHash()
	h.SetBytes([]byte(s))

	return h
}

// BytesToHash sets b to hash.
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

// Cross Type Routines

// BigToHash converts a big int to a Hash type
func BigToHash(b *big.Int) Hash {
	return BytesToHash(b.Bytes())
}

// Hash Type Routines

// IsValid verifies whether the hash is correct
func (h Hash) IsValid() bool {
	if len(h) != HashLength {
		return false
	}
	return true
}

// Hex converts a Hash to it's hex representation
func (h Hash) Hex() string {

	return ToHex(h[:])
}

// Big converts a Hash to a big int
func (h Hash) Big() *big.Int {
	return new(big.Int).SetBytes(h[:])
}

// Bytes convert a fixed 32 Hash to a byte slice
func (h Hash) Bytes() []byte {
	return h[:]
}

// String implements Stringer interface used for debugging purposes
func (h Hash) String() string {
	return h.Hex()
}

// SetBytes fills the hash with the value of b
func (h *Hash) SetBytes(b []byte) {
	if len(b) > HashLength {
		// crop b
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

// IsEqual compares two hashes
func (h *Hash) IsEqual(hash *Hash) bool {
	if h == nil && hash == nil {
		return true // comparing empty hashes
	}

	if h == nil || hash == nil {
		return false // comparing an empty hash to a non empty hash
	}

	return *h == *hash
}

// Blake2Hash calculates and returns the hash of the input data as a Hash type
func Blake2Hash(data []byte) Hash {
	var h Hash
	b := blake2b.Sum256(data)
	h.SetBytes(b[:])
	// h is a fixed [32]byte slice we return a view instead to ensure compatibility with other functions
	return h
}

// Blake2 returns the Blake2b hash as bytes
func Blake2(data []byte) []byte {

	h := Blake2Hash(data)
	return h.Bytes()
}

// Blake160 calculates and returns 160 bit hash of the input data this is only needed for shorter hashes
func Blake160(data []byte) []byte {

	hasher, err := blake2b.New(20, nil)

	if err != nil {
		return nil
	}
	// reset
	hasher.Reset()
	// write data to hasher
	hasher.Write(data)
	// sum
	h := hasher.Sum(nil)

	return h[:]
}
