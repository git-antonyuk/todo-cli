package cmd

import (
	"fmt"
	"strconv"
	"todo-console/cmd/tasks"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete <taskid>",
	Short: "Complete task by id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing task id: %s.\n\n", err)
			return
		}
		err = tasks.Complete(id)
		if err != nil {
			fmt.Printf("Error complete task id \"%s\": %s.\n\n", strconv.FormatInt(id, 10), err)
			return
		}
		fmt.Printf("Task id \"%s\" completed.\n\n", strconv.FormatInt(id, 10))
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
