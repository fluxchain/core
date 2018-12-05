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

func (b *Blockchain) HasGenesis() bool {
	genesis, err := storage.GetBlockByHeight(0)
	if err != nil {
		panic(err)
	}

	if genesis == nil {
		return false
	}

	return true
}

// Walks over all the blocks in storage, validating them, gathering some statistics
// and setting the tip to the appropriate state.
func (b *Blockchain) Hydrate() error {
	storage.WalkBlocks(func(currentBlock *block.Block) error {

		if b.Tip == nil || currentBlock.Header.Height > b.Tip.Header.Height {
			b.Tip = currentBlock
		}

		return nil
	})

	return nil
}

// Adds a block to the chain if it passes some validation.
func (b *Blockchain) AddBlock(currentBlock *block.Block) error {
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
	// ensure the prevhash this block is referring to is actually the tip
	// with exception of the genesis block.
	prevHash := hex.EncodeToString(currentBlock.Header.PrevHash)
	if currentBlock.Header.Height != 0 &&
		hex.EncodeToString(b.Tip.Header.Hash) != prevHash {
		return fmt.Errorf(
			"block has parent that isn't the current tip block: %v",
			currentBlock.Header.Hash)
	}

	// ensure the height is that of the previous block +1 with the exception of the genesis block.
	if b.Tip != nil && currentBlock.Header.Height != (b.Tip.Header.Height+1) {
		return fmt.Errorf(
			"block %v being added has an unexpected height. expected %v, got %v",
			currentBlock.Header.Hash,
			b.Tip.Header.Height+1,
			currentBlock.Header.Height)
	}

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
