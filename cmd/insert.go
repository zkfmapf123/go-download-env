package cmd

import (
	"log"
	"sync"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/go-download-env/internal/aws"
	"github.com/zkfmapf123/go-download-env/internal/filesystem"
	"github.com/zkfmapf123/go-download-env/internal/interaction"
)

var (
	cliList = []string{"env file upload", "cli"}
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

			// 3. select cli or env file upload
			v := interaction.SelectBox("Select cli or env file upload", cliList)

			if v == "cli" {
				updateKeyValue(awsParams, env, project)
				return 
			}

			uploadEnvFile(awsParams, env, project)
		},
	}
)

func uploadEnvFile(awsParams aws.AWSEnvParmas, env string, project string) {

	files, err := filesystem.GetEnvFilesCurrentDir()
	if err != nil {
		log.Fatalln(err)
	}

	if len(files) == 0 {
		log.Println("no env file found")
		return 
	}

	envFile := interaction.SelectBox("Select env file", files)
	envMap, err := filesystem.EnvFileToMap(envFile)
	if err != nil {
		log.Fatalln(err)
	}

	wg := sync.WaitGroup{}
	for k , v := range envMap{
		wg.Add(1)
	
		go func(key string, value string, env string, project string) {
			defer wg.Done()

			err := awsParams.PutSecretManager(key, value, env, project)
			if err != nil {
				panic(err)
			}
		}(k ,v, env, project)
	}
	
	wg.Wait()
	
}

func updateKeyValue(awsParams aws.AWSEnvParmas, env string, project string) {

}

func init() {
	rootCmd.AddCommand(insertCmd)
}
