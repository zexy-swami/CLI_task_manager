package db

import (
	"github.com/boltdb/bolt"
	"github.com/zexy-swami/CLI_task_manager/pkg/conversions"
)

type Task struct {
	Key  uint64
	Task string
}

func ReadTasks() ([]Task, error) {
	var readTasks []Task

	err := database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(tasksBucket)
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			readTasks = append(readTasks, Task{
				Key:  conversions.Btoi(k),
				Task: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return readTasks, nil
}
