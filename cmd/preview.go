package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/go-download-env/internal/aws"
	"github.com/zkfmapf123/go-download-env/internal/filesystem"
	"github.com/zkfmapf123/go-download-env/internal/interaction"
)

var (
	previewCmd = &cobra.Command{
		Use: "preview",
		Short: "preview is a command for previewing AWS credentials",
		Run: func(cmd *cobra.Command, args []string) {
			interaction.Clear()
			awsParams, fs := aws.MustNewAWS(), filesystem.NewFS()

			value, err := awsParams.GetParameter()
			if err != nil {
				log.Fatalln(err)
			}

			if len(value.Envs) == 0 && len(value.Projects) == 0 {
				log.Println("Not Exists env , project")
			}	

			fs.Dashboard([]string{"env"}, [][]string{value.Envs})
			fs.Dashboard([]string{"proejct"}, [][]string{value.Projects})

			interaction.PressEnter("Press Enter")
		},
	}
)

func init() {
	rootCmd.AddCommand(previewCmd)
}