package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/go-download-env/internal/aws"
	"github.com/zkfmapf123/go-download-env/internal/interaction"
	gojsmodule "github.com/zkfmapf123/go-js-utils"
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

			s3, isExist := awsParams.IsExistBucket()
			if !isExist {
				log.Fatalf("%s bucket is not exist", awsParams.GetS3Bucket())
			}

			pjtFullStr := []string{}
			for _, obj := range s3.Contents {
				key := obj.Key
				pjtFullStr = append(pjtFullStr, *key)
			}

			m := GetSelecProjects(pjtFullStr)

			// 1. select team
			team := gojsmodule.OKeys(m)
			selectTeam := interaction.SelectBox("[1] select Team ", team)

			// except common
			if selectTeam == "common" {

				selectKey := interaction.InputText("[2] key name")
				selectValue := interaction.InputText("[3] value")
				MustUpdateTask(awsParams, "common", selectKey, selectValue)

			}else{

				// 2. select project
				teamEnvs := map[string][]string{}
				for _, value := range m[selectTeam] {
					for k ,v := range value {
						teamEnvs[k] = append(teamEnvs[k], v)
					}
				}
	
				pjt := gojsmodule.OKeys(teamEnvs)
				selectPjt := interaction.SelectBox("[2] select Project ", pjt)
				
				// 3. select env
				envs := teamEnvs[selectPjt]
				envMap := map[string]string{}
				for _, env := range envs{
					envMap[env] = env
				}

				envs = gojsmodule.OKeys(envMap)
				selectEnv := interaction.SelectBox("[3] select Env ", envs)
	
				// 4. key name 
				selectKey := interaction.InputText("[4] key name")
	
				// 5. value 
				selectValue := interaction.InputText("[5] value")
				
				// 6. Update S3 / SSM Upload
				fullPath := fmt.Sprintf("%s/%s/%s", selectTeam, selectPjt, selectEnv)
				MustUpdateTask(awsParams, fullPath, selectKey,selectValue)
			}

			interaction.PressEnter("Success Insert > Press Enter")
		},
	}
)

func MustUpdateTask(awsParams aws.AWSEnvParmas,fullPath, key, value string) {
	err := awsParams.PutObject(fullPath, key, value)
	if err != nil {
		log.Fatalln(err)
	}

	err = awsParams.CreateParameter(fmt.Sprintf("/%s/%s", fullPath, key), value)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetSelecProjects(vs []string) map[string][]map[string]string{

	m := map[string][]map[string]string{}

	for _, v := range vs {
		split := strings.Split(v, "/")

		team := split[0]

		if len(split) < 3 {
			m[team] = []map[string]string{}
			continue
		}

		pjt := split[1]
		env := split[2]

		_m := map[string]string{
			pjt : env,
		}

		m[team]=append(m[team], _m)
	}

	return m
}

func init() {
	rootCmd.AddCommand(insertCmd)
}
