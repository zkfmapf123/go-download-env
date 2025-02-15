package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spf13/viper"
)

type AWSEnvParmas struct {
	ssmParameterClient *ssm.Client
	secretManagerClient *secretsmanager.Client
	iamClient *iam.Client
	stsClient *sts.Client
	s3Client *s3.Client
	
	profile string
	region string
	role string // admin, developer, readonly
	s3Bucket string 
}

func MustNewAWS() AWSEnvParmas {
	
	profile, region, s3Bucket := viper.GetString("profile"), viper.GetString("region"), viper.GetString("s3")
	
	cfg,err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)

	if err != nil {
		log.Fatal(err)
	}

	awsConfig := AWSEnvParmas{
		ssmParameterClient: ssm.NewFromConfig(cfg),
		secretManagerClient: secretsmanager.NewFromConfig(cfg),
		iamClient: iam.NewFromConfig(cfg),
		stsClient: sts.NewFromConfig(cfg),
		s3Client: s3.NewFromConfig(cfg),
		profile: profile,
		region: region,
		s3Bucket: s3Bucket,
	}

	role, err := awsConfig.GetUser()
	if err != nil {
		log.Fatal(err)
	}

	awsConfig.role = role
	return awsConfig
}

func (a *AWSEnvParmas) GetProfile() string {
	return a.profile
}

func (a *AWSEnvParmas) GetRegion() string {
	return a.region
}

func (a *AWSEnvParmas) GetS3Bucket() string {
	return a.s3Bucket
}

// 추후 활용할 예정
func (a *AWSEnvParmas) GetRole() string {
	return a.role
}

// 추후 활용할 예정
func (a *AWSEnvParmas) FatalErrorDeveloper() {

	if a.role == "developer" {
		log.Fatalln("You are not admin !!")
	}
}

