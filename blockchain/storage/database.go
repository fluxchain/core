package storage

import (
	"github.com/boltdb/bolt"
)

const (
	BLOCK_BUCKET        = "blocks"
	BLOCK_HEIGHT_BUCKET = "blocks_height"
	TRANSACTION_BUCKET  = "transactions"
)

var db *bolt.DB

func OpenDatabase(path string) error {
	var err error
	db, err = bolt.Open(path, 0600, nil)

	return err
}

func CloseDatabase() {
	db.Close()
}

func Migrate() error {
	var err error
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BLOCK_BUCKET))
		return err
	})
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BLOCK_HEIGHT_BUCKET))
		return err
	})
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(TRANSACTION_BUCKET))
		return err
	})

	return err
}
