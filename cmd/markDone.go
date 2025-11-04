package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var markDoneCli = &cobra.Command{
	Use:  "mark-done",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var id, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("An error occurred while converting to int")
		}
		fmt.Println(MarkDone(id))
	},
}

func init() {
	rootCmd.AddCommand(markDoneCli)
}
