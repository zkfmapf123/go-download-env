package aws

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func (a *AWSEnvParmas) CreateSecretManager(name string) error {

	_, err := a.secretManagerClient.CreateSecret(context.TODO(), &secretsmanager.CreateSecretInput{
		Name: aws.String(name),
	})

	if err != nil {

		if strings.Contains(err.Error(), "already exists") {
			return nil
		}

		return err
	}

	return nil
}

func (a *AWSEnvParmas) PutSecretManager(name string, key string, value string) {

}

func (a *AWSEnvParmas) DeleteSecretManager(name string, key string) {

}

func (a *AWSEnvParmas) GetSecretManager(name string) error {

	_, err := a.secretManagerClient.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	})

	// ResourceNotFoundException
	if err != nil {
		return err
	}

	return nil

}