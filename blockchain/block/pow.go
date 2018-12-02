package block

import (
	"bytes"

	c "github.com/fluxchain/core/crypto"
)

// Generates a proof-of-work by simply generating a blockhash and checking if
// the first 4 characters are all zeroes. Need to rework this.
func (h *Header) GeneratePOW() []byte {
	var err error
	var hash c.Hash

	for {
		hash, err = h.CalculateHash()
		if err != nil {
			panic(err)
		}

		hashStr := hash.String()

		if hashStr[:4] == "0000" {
			break
		}

		h.Nonce += 1
	}

	return hash
}

// Checks if the resulting block hash generated by local calculations matches
// the blockheader hash.
func (h *Header) ValidatePOW() (bool, error) {
	headerHash := h.Hash
	calculatedHash, err := h.CalculateHash()

	if err != nil {
		return false, err
	}

	return bytes.Equal(headerHash, calculatedHash), nil
}
