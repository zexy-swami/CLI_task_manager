package commands

import "github.com/spf13/cobra"

func init() {
	RootCommand.AddCommand(AddCommand)
	RootCommand.AddCommand(DoCommand)
	RootCommand.AddCommand(ListCommand)
}

var RootCommand = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
