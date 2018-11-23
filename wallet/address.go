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

package wallet

// Flux addresses are zbase32 encoded
// The Operation is done like this :
// Hash := Blake2(PublicKey)
// ShortHash := Blake160(Hash)
// Checksum := 4 last bytes of Blake2(ShortHash)
// zbase32(Mainnet||ShortHash||Checksum)

import (
	"github.com/fluxchain/core/crypto"
	"github.com/tv42/zbase32"
)

// P2PKH transactions are paid to the shorthash
// Meaning that only he who owns a publickey that Blake160(Blake2(PK)) can use it

var (
	// Mainnet is the identifier for mainnet addresses
	Mainnet = []byte{0x29, 0x7a} // starts with fff
	// Testnet is the identifier for testnet addresses
	Testnet = []byte{0xfb, 0xef} // starts with 9xz

)

func genAddr(publickey crypto.PublicKey) []byte {

	pkhash := crypto.Blake2(publickey[:])
	shorthash := crypto.Blake160(pkhash)
	checksum := crypto.Blake2(shorthash)
	checksum = checksum[len(checksum)-4:]

	baddr := append(Testnet, shorthash...)
	baddr = append(baddr, checksum...)

	var out []byte

	zbase32.Encode(out, baddr)

	return out
}

// GenAddr generates a new Flux address
func GenAddr(publickey []byte) string {

	addr := genAddr(publickey)

	return string(addr)

}
