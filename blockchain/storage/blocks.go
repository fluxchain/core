package storage

import (
	"encoding/json"

	"github.com/fluxchain/core/blockchain/block"
	c "github.com/fluxchain/core/crypto"
	bolt "go.etcd.io/bbolt"
)

type blockRepository struct{}

func (r *blockRepository) Store(b *block.Block) error {
	serialized, err := serializeBlock(b)
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_BUCKET))
		err := bucket.Put([]byte(b.Header.Hash), serialized)

		return err
	})

	return err
}

func (r *blockRepository) Get(hash c.Hash) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCK_BUCKET))
		b.Get(hash)
		return nil
	})
}

func serializeBlock(b *block.Block) ([]byte, error) {
	return json.Marshal(b)
}

var BlockRepository *blockRepository
