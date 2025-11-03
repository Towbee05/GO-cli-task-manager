package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listTaskCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"list"},
	Short:   "View all tasks",
	Long:    "Get a list of all added tasks",
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		filter := ""
		if len(args) > 0 {
			filter = args[0]
		}
		fmt.Println(List(filter))
	},
}

func init() {
	rootCmd.AddCommand(listTaskCmd)
}
