package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/spf13/viper"
)

type AWSEnvParmas struct {
	ssmParameterClient *ssm.Client
	secretManagerClient *secretsmanager.Client
	
	Profile string
	Region string
}

func MustNewAWS() AWSEnvParmas {
	
	profile, region := viper.GetString("profile"), viper.GetString("region")
	
	cfg,err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)

	if err != nil {
		log.Fatal(err)
	}

	return AWSEnvParmas{
		ssmParameterClient: ssm.NewFromConfig(cfg),
		secretManagerClient: secretsmanager.NewFromConfig(cfg),
		Profile : profile,
		Region : region,
	}
}