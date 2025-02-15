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
			fs := filesystem.NewFS()

			projectSchema, err := filesystem.GetYamlFileData()
			if err != nil {
				log.Fatal(err)
			}

			err = awsParams.UpdateS3Architecture(projectSchema)
			if err != nil {
				panic(err)
			}

			headers := []string{"team", "project", "env"}
			body := [][]string{}

			if projectSchema.IsUseCommonEnvironments {
				body = append(body, []string{"common"})
			}

			for team, value := range projectSchema.Projects {
				for _, project := range value {
					for _, env := range projectSchema.Envs {
						body = append(body, []string{team, project, env})
					}
				}
			}

			fs.Dashboard(headers, body)
			interaction.PressEnter("Success Setting > Press Enter")
		},
	}
)

func init() {
	rootCmd.AddCommand(settingCmd)
}