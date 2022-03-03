package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zexy-swami/CLI_task_manager/internal/db"
)

var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run:   ListAction,
}

func ListAction(command *cobra.Command, args []string) {
	tasks, err := db.ReadTasks()
	if err != nil {
		fmt.Printf("Error occurred while reading tasks. Error: %s\n", err.Error())
		return
	}

	if len(tasks) == 0 {
		fmt.Println("You have no tasks to complete")
		return
	}

	fmt.Println("Tasks:")
	for taskIndex, task := range tasks {
		fmt.Printf("%d. %s\n", taskIndex+1, task.Task)
	}
}
