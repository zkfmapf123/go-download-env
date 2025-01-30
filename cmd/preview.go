package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	previewCmd = &cobra.Command{
		Use: "preview",
		Short: "preview is a command for previewing AWS credentials",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("preview")
		},
	}
)

func init() {
	rootCmd.AddCommand(previewCmd)
}