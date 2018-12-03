package storage

import (
	"encoding/json"

	"github.com/fluxchain/core/blockchain/block"
	c "github.com/fluxchain/core/crypto"
	bolt "go.etcd.io/bbolt"
)

func StoreBlock(b *block.Block) error {
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

func GetBlock(hash c.Hash) (*block.Block, error) {
	var result *block.Block
	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(BLOCK_BUCKET))

		data := b.Get(hash)
		result, err = deserializeBlock(data)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func LastBlock() (*block.Block, error) {
	var result *block.Block
	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(BLOCK_BUCKET))

		c := b.Cursor()
		_, data := c.Last()

		result, err = deserializeBlock(data)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func serializeBlock(b *block.Block) ([]byte, error) {
	return json.Marshal(b)
}

func deserializeBlock(data []byte) (*block.Block, error) {
	var result *block.Block

	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
