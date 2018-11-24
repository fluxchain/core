package blockchain

import (
	"github.com/cbergoon/merkletree"
)

// Body contains all transactions within a block
type Body []*Transaction

// MerkleTree creates a merlketree from the transaction hashes in the block.
func (b *Body) MerkleTree() *merkletree.MerkleTree {
	var list []merkletree.Content

	for _, transaction := range *b {
		list = append(list, transaction)
	}

	t, _ := merkletree.NewTree(list)

	return t
}
