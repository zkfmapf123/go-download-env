package aws

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func (a *AWSEnvParmas) CreateSecretManager(name string) error {

	log.Println("generate secret manager", name)
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

func (a *AWSEnvParmas) PutSecretManager(key, value string) error {

    // 1. 먼저 시크릿이 존재하는지 확인
    _, err := a.secretManagerClient.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
        SecretId: aws.String(key),
    })

    if err != nil {
        // 시크릿이 없으면 생성
        if strings.Contains(err.Error(), "ResourceNotFoundException") {
            _, err = a.secretManagerClient.CreateSecret(context.TODO(), &secretsmanager.CreateSecretInput{
                Name:         aws.String(key),
                SecretString: aws.String(value),
            })

            return err
        }
        return err
    }

    _, err = a.secretManagerClient.PutSecretValue(context.TODO(), &secretsmanager.PutSecretValueInput{
        SecretId:     aws.String(key),
        SecretString: aws.String(value),
    })

    return err
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