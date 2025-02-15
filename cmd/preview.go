package cmd

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/go-download-env/internal/aws"
	"github.com/zkfmapf123/go-download-env/internal/filesystem"
	"github.com/zkfmapf123/go-download-env/internal/interaction"
	gojsmodule "github.com/zkfmapf123/go-js-utils"
)

var (
	previewCmd = &cobra.Command{
		Use: "preview",
		Short: "preview is a command for previewing AWS credentials",
		Run: func(cmd *cobra.Command, args []string) {
			interaction.Clear()
			awsParams, fs := aws.MustNewAWS(), filesystem.NewFS()

			s3, isExists := awsParams.IsExistBucket() 
			if !isExists {
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

			if selectTeam == "common" {

			}

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
			selectEnv := interaction.SelectBox("[3] select Env", envs)

			prefix := fmt.Sprintf("%s/%s/%s", selectTeam, selectPjt, selectEnv)
			ssmKeys := []string{}
			for _, obj := range s3.Contents {
				key := *obj.Key
				
				if strings.Contains(key, prefix) && len(strings.Split(key,"/")) > 3  && strings.Split(key, "/")[3] != ""{
					ssmKeys = append(ssmKeys, key)			
				}
			}

			var wg sync.WaitGroup
			ssmMap := make(map[string]string)
			var mu sync.Mutex

			for _ ,ssmKey := range ssmKeys {
				wg.Add(1)
				go func(ssmKey string) {
					defer wg.Done()
					
					ssmValue, err := awsParams.GetParameter(fmt.Sprintf("/%s",ssmKey))
					if err != nil {
						log.Fatalf("Error getting parameter %s: %v", ssmKey, err)
						return
					}
					
					mu.Lock()
					ssmMap[ssmKey] = ssmValue
					mu.Unlock()
				}(ssmKey)
			}
				
			wg.Wait()

			body := [][]string{}
			for k, v := range ssmMap {
				splitKey := strings.Split(k, "/") 
				lastPart := splitKey[len(splitKey)-1] 
				body = append(body, []string{k, lastPart,v})
			}
			
			fs.Dashboard([]string{"full path","key","value"},body)
		},
	}
)

func init() {
	rootCmd.AddCommand(previewCmd)
}