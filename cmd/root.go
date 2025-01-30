package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zkfmapf123/go-download-env/internal/aws"
	"github.com/zkfmapf123/go-download-env/internal/filesystem"
	"github.com/zkfmapf123/go-download-env/internal/interaction"
)

var (
	_defaultProfile = "default"
	_defaultRegion = "ap-northeast-2"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-envs",
		Short: "go-envs is a tool for managing AWS credentials",
		Run: func(cmd *cobra.Command, args []string) {
			interaction.Clear()
			fs := filesystem.NewFS()
			
			awsParams := aws.MustNewAWS()
			
			fs.Dashboard([]string{"profile", "region", "role"}, [][]string{{awsParams.GetProfile(), awsParams.GetRegion(), awsParams.GetRole()}})
			
			_, err := awsParams.GetParameter()
			if err != nil {
				panic(err)
			}
		},
	}
)

func initial() {
	// init
	cobra.OnInitialize(func() {

		if viper.GetString("profile") == "" {
			viper.Set("profile", _defaultProfile)
		}

		if viper.GetString("region") == "" {
			viper.Set("region", _defaultRegion)
		}
	})

	rootCmd.PersistentFlags().StringP("profile", "p", _defaultProfile, "[Optional] AWS profile name")
	rootCmd.PersistentFlags().StringP("region", "r", _defaultRegion, "[Optional] AWS region")

	err := viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))
	if err != nil {
		log.Fatalln(err)
	}

	err = viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region"))
	if err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	initial()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}