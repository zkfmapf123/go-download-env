package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const (
	KEY = "/goenvs/settings"
)

func (a *AWSEnvParmas) GetParameter(key string) (string, error) {

	resp, err := a.ssmParameterClient.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: aws.String(key),
	})

	if err != nil {
		return "", err
	}

	return *resp.Parameter.Value, nil

}

func (a *AWSEnvParmas) CreateParameter(key string, value string) error {
	

	_, err := a.ssmParameterClient.PutParameter(context.TODO(), &ssm.PutParameterInput{
		Name: aws.String(key),
		Value: aws.String(value),
		DataType: aws.String("text"),
		Description: aws.String("Environment Settings"),
		Overwrite: aws.Bool(true),
		Type: types.ParameterTypeString,
	})

	if err != nil {
		return err
	}

	return nil
}