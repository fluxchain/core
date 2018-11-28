package block

import (
	"fmt"

	"github.com/cbergoon/merkletree"
	"github.com/fluxchain/core/blockchain/transaction"
)

type Body []*transaction.Transaction

// Creates a merlketree from the transaction hashes in the block.
func (b *Body) CalculateMerkle() *merkletree.MerkleTree {
	var list []merkletree.Content

	for _, transaction := range *b {
		list = append(list, transaction)
	}

	t, _ := merkletree.NewTree(list)

	return t
}

// Adds the transaction to the block body, after a given set of validation
// rules.
func (b *Body) AddTransaction(transaction *transaction.Transaction) error {
	if err := b.ValidateTransaction(transaction); err != nil {
		return err
	}

	*b = append(*b, transaction)

	return nil
}

func (b *Body) ValidateTransaction(tx *transaction.Transaction) error {
	for _, output := range tx.Outputs {
		if !output.Recipient.Valid() {
			return fmt.Errorf("Transaction address is invalid %v", output.Recipient.String)
		}
	}

	return nil
}

func NewBody() *Body {
	return &Body{}
}
