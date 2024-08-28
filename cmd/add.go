package cmd

import (
	"fmt"
	"todo-console/cmd/tasks"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "Add new task",
	Long:  "Carry out addition operation on 2 numbers",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := tasks.Add(args[0])
		if err != nil {
			fmt.Printf("Error adding new task: %s.\n\n", err)
			return
		}
		fmt.Printf("Task with description \"%s\" added.\n\n", args[0])
	},
}

func init() {

	rootCmd.AddCommand(addCmd)
}
