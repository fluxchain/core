package blockchain

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Blockchain struct {
	Blocks []*Block `json:"blocks"`
	Tip    *Block   `json:"-"`
}

// Adds a block to the chain if it passes some validation.
func (b *Blockchain) AddBlock(block *Block) error {
	// ensure the prevhash this block is referring to is actually the tip
	// with exception of ofcourse the genesis block.
	prevHash := hex.EncodeToString(block.Header.PrevHash)
	if block.Header.Height != 0 &&
		hex.EncodeToString(b.Tip.Header.Hash) != prevHash {
		return fmt.Errorf(
			"block has parent that isn't the current tip block: %v",
			block.Header.Hash)
	}

	// ensure the height is that of the previous block +1 with the obvious
	// exception of the genesis block
	if b.Tip != nil && block.Header.Height != (b.Tip.Header.Height+1) {
		return fmt.Errorf(
			"block %v being added has an unexpected height. expected %v, got %v",
			block.Header.Hash,
			b.Tip.Header.Height+1,
			block.Header.Height)
	}

	// Validate the block itself
	if err := b.ValidateBlock(block); err != nil {
		return err
	}

	log.Printf("adding block: %v", block)
	b.Blocks = append(b.Blocks, block)
	b.Tip = block

	return nil
}

// Validates the to-be-added block, currently only checks the validity of the
// PoW.
func (b *Blockchain) ValidateBlock(block *Block) error {
	// check if PoW checks out
	if !block.Header.ValidatePOW() {
		return fmt.Errorf("POW seems invalid for block %v",
			block.Header.Hash.String())
	}

	return nil
}

// Serializes the chain to JSON and writes it to the passed in writer.
func (b *Blockchain) Serialize(w io.Writer) error {
	enc := json.NewEncoder(w)

	return enc.Encode(b)
}

// Creates an instance of a new chain.
func NewBlockchain() *Blockchain {
	return &Blockchain{}
}
