package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	downloadCmd = &cobra.Command{
		Use: "download",
		Short: "download is a command for downloading AWS credentials",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("download")
		},
	}
)

func init() {
	rootCmd.AddCommand(downloadCmd)
}
