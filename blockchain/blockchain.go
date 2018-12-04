package blockchain

import (
	"encoding/hex"
	"fmt"

	"github.com/fluxchain/core/blockchain/block"
	"github.com/fluxchain/core/blockchain/storage"
	"github.com/fluxchain/core/consensus"
	"github.com/fluxchain/core/parameters"
)

type Blockchain struct {
	Tip *block.Block
}

func (b *Blockchain) HasGenesis() (bool, error) {
	return storage.HasBlockHeight(0)
}

func (b *Blockchain) Hydrate() error {
	storage.WalkBlocks(func(currentBlock *block.Block) error {
		// ugly hack until I get the data sorted
		if currentBlock.Header.Height > b.Tip.Header.Height {
			b.Tip = currentBlock
		}

		return nil
	})

	return nil
}

// Adds a block to the chain if it passes some validation.
func (b *Blockchain) AddBlock(currentBlock *block.Block) error {
	// ensure the prevhash this block is referring to is actually the tip
	// with exception of ofcourse the genesis block.
	prevHash := hex.EncodeToString(currentBlock.Header.PrevHash)
	if currentBlock.Header.Height != 0 &&
		hex.EncodeToString(b.Tip.Header.Hash) != prevHash {
		return fmt.Errorf(
			"block has parent that isn't the current tip block: %v",
			currentBlock.Header.Hash)
	}

	// ensure the height is that of the previous block +1 with the obvious
	// exception of the genesis block
	if b.Tip != nil && currentBlock.Header.Height != (b.Tip.Header.Height+1) {
		return fmt.Errorf(
			"block %v being added has an unexpected height. expected %v, got %v",
			currentBlock.Header.Hash,
			b.Tip.Header.Height+1,
			currentBlock.Header.Height)
	}

	// Validate the block itself
	if err := b.ValidateBlock(currentBlock); err != nil {
		return err
	}

	if err := storage.StoreBlock(currentBlock); err != nil {
		return err
	}

	b.Tip = currentBlock
	fmt.Printf("Updated tip %v\n", b.Tip)

	return nil
}

// Validates the to-be-added block, currently only checks the validity of the
// PoW.
func (b *Blockchain) ValidateBlock(currentBlock *block.Block) error {
	// check if PoW checks out
	powValid, err := consensus.ValidatePOW(currentBlock.Header, parameters.Current().MinimumPoW)
	if err != nil {
		return err
	}

	if !powValid {
		return fmt.Errorf("POW seems invalid for block %v",
			currentBlock.Header.Hash.String())
	}

	return nil
}

// Creates an instance of a new chain.
func NewBlockchain() *Blockchain {
	return &Blockchain{}
}
