package blockchain

import (
	"encoding/hex"
	"fmt"

	"github.com/fluxchain/core/blockchain/block"
	"github.com/fluxchain/core/blockchain/storage"
	"github.com/fluxchain/core/consensus"
	"github.com/fluxchain/core/parameters"
	"github.com/sirupsen/logrus"
)

type Blockchain struct {
	Tip *block.Block
}

// looks up the genesis block and return whether or not it can be found.
func (b *Blockchain) HasGenesis() (bool, error) {
	genesis, err := storage.GetBlockByHeight(0)
	if err != nil {
		return false, err
	}
	if genesis == nil {
		return false, nil
	}

	return true, nil
}

// sets the tip of the blockchain to the specified block.
func (b *Blockchain) SetTip(tip *block.Block) {
	logrus.WithFields(logrus.Fields{
		"height": tip.Header.Height,
		"hash":   tip.Header.Hash,
	}).Debug("setting tip")

	b.Tip = tip
}

// Walks over all the blocks in storage, validating them, gathering some statistics
// and setting the tip to the appropriate state.
func (b *Blockchain) Hydrate() error {
	return storage.WalkBlocks(func(currentBlock *block.Block) error {
		logrus.WithFields(logrus.Fields{
			"height": currentBlock.Header.Height,
			"hash":   currentBlock.Header.Hash,
		}).Trace("hydrate looping over local block")

		if err := b.ValidateBlock(currentBlock); err != nil {
			return err
		}

		if b.Tip == nil || currentBlock.Header.Height > b.Tip.Header.Height {
			b.SetTip(currentBlock)
		}

		return nil
	})
}

// Adds a block to the chain if it passes some validation.
func (b *Blockchain) AddGenesisBlock(currentBlock *block.Block) error {
	if err := b.ValidateBlock(currentBlock); err != nil {
		return err
	}

	if err := storage.StoreBlock(currentBlock); err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"hash":   currentBlock.Header.Hash,
		"height": currentBlock.Header.Height,
	}).Info("added genesis block")

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

	logrus.WithFields(logrus.Fields{
		"hash":   currentBlock.Header.Hash,
		"height": currentBlock.Header.Height,
	}).Info("added block")

	b.SetTip(currentBlock)

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
