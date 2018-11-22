package wallet

import (
	"bytes"
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"

	"github.com/m0t0k1ch1/base58"
)

const CHECKSUM_LENGTH = 4

var b58 = base58.NewBitcoinBase58()

type Address struct {
	PublicKey []byte `json:"-"`
	Checksum  []byte `json:"-"`
	String    string `json:"string"`
}

func (a *Address) Valid() bool {
	decoded, err := b58.DecodeString(a.String)
	if err != nil {
		log.Panic(err)
	}

	publicHash := decoded[:len(decoded)-CHECKSUM_LENGTH]
	actualChecksum := grabChecksum(decoded)
	expectedChecksum := calculateChecksum(publicHash)

	return bytes.Compare(actualChecksum, expectedChecksum) == 0
}

// Create a new address instance from a public key,
func NewAddressFromPublicKey(publicKey []byte) *Address {
	sha := sha256.Sum256(publicKey)

	hash := ripemd160.New()
	_, err := hash.Write(sha[:])
	if err != nil {
		log.Panic(err)
	}
	publicHash := hash.Sum(nil)

	checksum := calculateChecksum(publicHash)

	address, err := b58.EncodeToString(append(publicHash, checksum...))
	if err != nil {
		log.Panic(err)
	}

	return &Address{
		PublicKey: publicKey,
		String:    address}
}

// Create a new address instance from an address string
func NewAddressFromString(address string) *Address {
	return &Address{String: address}
}

func grabChecksum(a []byte) []byte {
	return a[len(a)-CHECKSUM_LENGTH:]
}

func calculateChecksum(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])

	return second[:CHECKSUM_LENGTH]
}
