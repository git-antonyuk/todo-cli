package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "tasks is a cli tool",
	Long: `
	Tasks is a cli tool for creation todo list

	Example of usage: 
	$ tasks add "My new task"
	$ tasks list
	$ tasks complete 
	`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
