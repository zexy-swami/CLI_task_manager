package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var (
	tasksBucket = []byte("Tasks")
	database    *bolt.DB
)

func InitDB(pathToDatabase string) error {
	var err error
	database, err = bolt.Open(pathToDatabase, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return database.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		return err
	})
}
