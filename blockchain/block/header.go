package block

import (
	"encoding/binary"
	"time"

	c "github.com/fluxchain/core/crypto"
	"golang.org/x/crypto/blake2b"
)

type Header struct {
	Height     uint32    `json:"height"`
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

	hash.Write(h.PrevHash)
	hash.Write(h.MerkleRoot)

	tsBuf := make([]byte, 8)
	binary.BigEndian.PutUint64(tsBuf, uint64(h.Height))
	hash.Write(tsBuf)

	tsBuf = make([]byte, 8)
	binary.BigEndian.PutUint64(tsBuf, uint64(h.Nonce))
	hash.Write(tsBuf)

	md := hash.Sum(nil)
	return md, nil
}
