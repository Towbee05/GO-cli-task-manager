package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteTaskCli = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "remove"},
	Short:   "Delete Task",
	Long:    "Delete task with specified ID",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var id, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("An error occurred while converting to int")
		}
		fmt.Println(DeleteTask(id))
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCli)
}
