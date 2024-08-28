package cmd

import (
	"fmt"
	"strconv"
	"todo-console/cmd/tasks"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <taskid>",
	Short: "Delete task by id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing task id: %s.\n\n", err)
			return
		}
		err = tasks.Delete(id)
		if err != nil {
			fmt.Printf("Error deleting task id \"%s\": %s.\n\n", strconv.FormatInt(id, 10), err)
			return
		}
		fmt.Printf("Task id \"%s\" deleted.\n\n", strconv.FormatInt(id, 10))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
