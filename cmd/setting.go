package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/go-download-env/internal/aws"
	"github.com/zkfmapf123/go-download-env/internal/filesystem"
	"github.com/zkfmapf123/go-download-env/internal/interaction"
)

var (
	settingCmd = &cobra.Command{
		Use:   "setting",
		Short: "setting is a tool for setting AWS credentials",
		Run: func(cmd *cobra.Command, args []string) {
			interaction.Clear()
			awsParams := aws.MustNewAWS()

			projectSchema, err := filesystem.GetYamlFileData()
			if err != nil {
				log.Fatal(err)
			}

			envParameter := aws.EnvParameterValues{
				Envs: projectSchema.Envs,
				Projects: projectSchema.Projects,
			}

			err = awsParams.CreateParameter(envParameter)
			if err != nil {
				panic(err)
			}

			interaction.PressEnter("Success Setting > Press Enter")
		},
	}
)

func init() {
	rootCmd.AddCommand(settingCmd)
}