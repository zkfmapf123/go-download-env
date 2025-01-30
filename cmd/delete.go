package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	deleteCmd = &cobra.Command{
		Use: "delete",
		Short: "delete is a command for deleting AWS credentials",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("delete")
		},
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}
