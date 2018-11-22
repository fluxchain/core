package blockchain

import (
	"fmt"
	"time"
)

type Block struct {
	Header *BlockHeader `json:"header"`
	Body   *BlockBody   `json:"transactions"`
}

func (b *Block) String() string {
	return fmt.Sprintf("[%v] height: %v",
		b.Header.Hash,
		b.Header.Height)
}

// Creates a new block, creates the merkletree and creates a valid PoW.
func NewBlock(prevblock *Block, timestamp time.Time, body *BlockBody) *Block {
	header := &BlockHeader{
		Height:    prevblock.Header.Height + 1,
		PrevHash:  prevblock.Header.Hash,
		Timestamp: timestamp,
	}

	// TODO make this less implicit to implementing methods.
	header.MerkleRoot = body.CalculateMerkle().MerkleRoot()
	header.Hash = header.GeneratePOW()

	return &Block{
		Header: header,
		Body:   body,
	}
}

// Creates a new genesis block, which doesn't have a previous block.
func NewGenesisBlock(timestamp time.Time, body *BlockBody) *Block {
	header := &BlockHeader{
		Height:    0,
		PrevHash:  []byte{},
		Timestamp: timestamp,
	}

	// TODO make this less implicit to implementing methods.
	header.MerkleRoot = body.CalculateMerkle().MerkleRoot()
	header.Hash = header.GeneratePOW()

	return &Block{
		Header: header,
		Body:   body,
	}
}
