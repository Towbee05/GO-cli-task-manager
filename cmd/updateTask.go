package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var updateTaskCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"change"},
	Short:   "Update Task",
	Long:    "Update your task, providing id and new name",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var id, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error while converting string")
		}
		fmt.Println(UpdateTask(id, args[1]))
	},
}

func init() {
	rootCmd.AddCommand(updateTaskCmd)
}
