package aws

import (
	"context"
	"fmt"
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

func (a *AWSEnvParmas) PutSecretManager(key, value, env, project string) error{
	err := a.GetSecretManager(fmt.Sprintf("%s-%s", env, project))

	// 리소스가 없다면 생성
	if err != nil {
		fmt.Println(err.Error())

		if strings.Contains(err.Error(), "ResourceNotFoundException") {
			err := a.CreateSecretManager(fmt.Sprintf("%s-%s", env, project))
			if err != nil {
				return err
			}
		}else {
			return err
		}
	}
	
	// put key, value 
	_ ,err = a.secretManagerClient.PutSecretValue(context.TODO(), &secretsmanager.PutSecretValueInput{
		SecretId: aws.String(fmt.Sprintf("%s-%s", env, project)),
		SecretString: aws.String(fmt.Sprintf("%s=%s", key, value)),
	})

	if err != nil {
		return err
	}

	return nil
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