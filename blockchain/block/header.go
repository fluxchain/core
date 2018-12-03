package block

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"time"

	c "github.com/fluxchain/core/crypto"
	"golang.org/x/crypto/blake2b"
)

type Header struct {
	Height     uint64    `json:"height"`
	Hash       c.Hash    `json:"hash"`
	PrevHash   c.Hash    `json:"prevhash"`
	Timestamp  time.Time `json:"timestamp"`
	MerkleRoot c.Hash    `json:"merkleroot"`
	Nonce      uint64    `json:"nonce"`
}

// Calculates the blockheader hash by concatting the previous blockhash, the
// merlkeroot, the height as a BE uint64 and the nonce as a BE uint64. And
// passing it through a round of SHA256
func (h Header) CalculateHash() (c.Hash, error) {
	hash, err := blake2b.New256(nil)
	if err != nil {
		return nil, err
	}

	serialized, err := h.SerializeForProof()
	if err != nil {
		return nil, err
	}
	hash.Write(serialized)

	md := hash.Sum(nil)
	return md, nil
}

func (h *Header) IncrementNonce() {
	h.Nonce = h.Nonce + 1
}

func (h Header) SerializeForProof() ([]byte, error) {
	var result bytes.Buffer
	buffer := bufio.NewWriter(&result)

	buffer.Write(h.PrevHash)
	buffer.Write(h.MerkleRoot)

	binary.Write(buffer, binary.BigEndian, uint64(h.Height))
	binary.Write(buffer, binary.BigEndian, uint64(h.Nonce))

	buffer.Flush()

	return result.Bytes(), nil
}

func NewHeader(prevHash c.Hash, height uint64, timestamp time.Time) *Header {
	return &Header{
		Height:    height,
		PrevHash:  prevHash,
		Timestamp: timestamp,
	}
}
