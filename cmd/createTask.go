package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createTaskCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create-task", "create"},
	Short:   "Create a new Task",
	Long:    "Add a new task or add to existing tasks",
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		description := ""
		if len(args) >= 1 {
			description = args[1]
		}
		fmt.Println(AddTask(task, description))
	},
}

func init() {
	rootCmd.AddCommand(createTaskCmd)
}
