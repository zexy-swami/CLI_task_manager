package db

import (
	"github.com/boltdb/bolt"
	"github.com/zexy-swami/CLI_task_manager/pkg/conversions"
)

func CreateTask(taskToCreate string) error {
	err := database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(tasksBucket)
		taskID, _ := bucket.NextSequence()
		taskKey := conversions.Itob(taskID)
		return bucket.Put(taskKey, []byte(taskToCreate))
	})

	return err
}
