package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/zexy-swami/CLI_task_manager/internal/db"
)

var DoCommand = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run:   DoAction,
}

func DoAction(command *cobra.Command, args []string) {
	taskIDs := make([]int, 0, len(args))
	for _, arg := range args {
		if taskID, err := strconv.Atoi(arg); err != nil {
			fmt.Printf("Failed to parse the argument: %s\n", arg)
		} else {
			taskIDs = append(taskIDs, taskID)
		}
	}

	tasks, err := db.ReadTasks()
	if err != nil {
		fmt.Printf("Error occurred while reading tasks. Error: %s\n", err.Error())
		return
	}

	for _, id := range taskIDs {
		if id <= 0 || id > len(tasks) {
			fmt.Printf("Invalid task number: %d\n", id)
			continue
		}

		task := tasks[id-1]

		if err := db.DeleteTask(task.Key); err != nil {
			fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
		} else {
			fmt.Printf("Marked \"%d\" as completed.\n", id)
		}
	}
}
