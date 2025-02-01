package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/go-download-env/internal/aws"
	"github.com/zkfmapf123/go-download-env/internal/interaction"
)

var (
	insertCmd = &cobra.Command{
		Use: "insert",
		Short: "insert is a command for inserting AWS credentials",
		Run: func(cmd *cobra.Command, args []string) {
			awsParams := aws.MustNewAWS()
			interaction.Clear()

			values, err := awsParams.GetParameter()
			if err != nil {
				log.Fatalln(err)
			}

			// 1. select envs
			env := interaction.SelectBox("Select Env", values.Envs)

			// 2. select projects
			project := interaction.SelectBox("Select Project", values.Projects)
			

			fmt.Println(env, project)
		},
	}
)

func init() {
	rootCmd.AddCommand(insertCmd)
}
