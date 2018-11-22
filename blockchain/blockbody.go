package blockchain

import (
	"fmt"

	"github.com/cbergoon/merkletree"
)

type BlockBody []*Transaction

// Creates a merlketree from the transaction hashes in the block.
func (b *BlockBody) CalculateMerkle() *merkletree.MerkleTree {
	var list []merkletree.Content

	for _, transaction := range *b {
		list = append(list, transaction)
	}

	t, _ := merkletree.NewTree(list)

	return t
}

// Adds the transaction to the block body, after a given set of validation
// rules.
func (b *BlockBody) AddTransaction(transaction *Transaction) error {
	if err := b.ValidateTransaction(transaction); err != nil {
		return err
	}

	*b = append(*b, transaction)

	return nil
}

func (b *BlockBody) ValidateTransaction(tx *Transaction) error {
	for _, output := range tx.Outputs {
		if !output.Recipient.Valid() {
			return fmt.Errorf("Transaction address is invalid %v", output.Recipient.String)
		}
	}

	return nil
}

func NewBlockBody() *BlockBody {
	return &BlockBody{}
}
