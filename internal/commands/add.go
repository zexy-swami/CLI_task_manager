package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zexy-swami/CLI_task_manager/internal/db"
)

var AddCommand = &cobra.Command{
	Use:   "add",
	Short: "Adds task to your task list",
	Run:   AddAction,
}

func AddAction(command *cobra.Command, args []string) {
	taskToAdd := strings.Join(args, " ")

	if err := db.CreateTask(taskToAdd); err != nil {
		fmt.Printf("Error occurred while creating task. Error: %s\n", err.Error())
		return
	}
	fmt.Printf("Added \"%s\" to your task list\n", taskToAdd)
}
