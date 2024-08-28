package cmd

import (
	"todo-console/cmd/tasks"

	"github.com/spf13/cobra"
)

var isAll bool
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks.List(isAll)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&isAll, "all", "a", false, "Show all columns")
	rootCmd.AddCommand(listCmd)
}
