package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use: "update",
		Short: "update is a command for updating AWS credentials",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("update")
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)
}