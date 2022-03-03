package db

import (
	"github.com/boltdb/bolt"
	"github.com/zexy-swami/CLI_task_manager/pkg/conversions"
)

func DeleteTask(taskKeyToDelete uint64) error {
	return database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(tasksBucket)
		return bucket.Delete(conversions.Itob(taskKeyToDelete))
	})
}
