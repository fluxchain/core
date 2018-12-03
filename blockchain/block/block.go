package block

import (
	"fmt"
	"time"
)

type Block struct {
	Header *Header `json:"header"`
	Body   *Body   `json:"transactions"`
}

func (b *Block) String() string {
	return fmt.Sprintf("[%v] height: %v",
		b.Header.Hash,
		b.Header.Height)
}

// Creates a new block, creates the merkletree and creates a valid PoW.
func NewBlock(prevBlock *Block, timestamp time.Time, body *Body) *Block {
	prevHeader := prevBlock.Header
	header := NewHeader(prevHeader.Hash, prevHeader.Height+1, timestamp)

	header.MerkleRoot = body.CalculateMerkle().MerkleRoot()

	return &Block{
		Header: header,
		Body:   body,
	}
}

// Creates a new genesis block, which doesn't have a previous block.
func NewGenesisBlock(timestamp time.Time, body *Body) *Block {
	header := NewHeader([]byte{}, 0, timestamp)

	header.MerkleRoot = body.CalculateMerkle().MerkleRoot()

	return &Block{
		Header: header,
		Body:   body,
	}
}
