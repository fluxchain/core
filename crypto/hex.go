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

package crypto

import (
	"encoding/base32"
	"encoding/hex"
	"errors"
)

var (
	errBadHex    = errors.New("string should contain only hexadecimal characters and be of even length")
	errBadBase32 = errors.New("string should be conform to base32 alphabet")
)

// ToHex encode bytes as Hex
func ToHex(b []byte) string {
	return hex.EncodeToString(b)
}

// ToBase32 encode bytes as Base32 string
func ToBase32(b []byte) string {
	return base32.StdEncoding.EncodeToString(b)
}

// FromHex decodes hex to bytes
func FromHex(s string) ([]byte, error) {
	b, err := hex.DecodeString(s)

	if err != nil {
		return nil, errBadHex
	}

	return b, nil
}

// FromBase32 decodes base32 encoded string to bytes
func FromBase32(s string) ([]byte, error) {
	b, err := base32.StdEncoding.DecodeString(s)

	if err != nil {
		return nil, errBadBase32
	}

	return b, nil
}

// Clone returns an exact copy of a byte slice
func Clone(b []byte) []byte {
	if b == nil {
		return nil
	}

	c := make([]byte, len(b))

	copy(c, b)

	return c

}

// Reset zeroes byte slice contents
func Reset(b []byte) {

	for i := range b {
		b[i] = 0
	}
}
