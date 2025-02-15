package aws

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const (
	KEY = "/goenvs/settings"
)

type EnvParameterValues struct {
	Envs []string `json:"envs"`
	Projects []string `json:"projects"`
}

func (a *AWSEnvParmas) GetParameter() (EnvParameterValues, error) {

	// resp, err := a.ssmParameterClient.GetParameter(context.TODO(), &ssm.GetParameterInput{
	// 	Name: aws.String(KEY),
	// })

	// if err != nil {

	// 	// Create
	// 	if strings.Contains(err.Error(), "ParameterNotFound") {
	
	// 		envValues := EnvParameterValues{
	// 			Envs: []string{},
	// 			Projects: []string{},
	// 		}

	// 		if err = a.CreateParameter(envValues); err != nil {
	// 			return EnvParameterValues{} ,err
	// 		}
	
	// 		return EnvParameterValues{}, nil
	// 	}
		
	// 	return EnvParameterValues{}, err
	// }

	// return mustStirngToJson(*resp.Parameter.Value), nil
	return EnvParameterValues{}, nil
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

func mustStirngToJson(value string) EnvParameterValues {
	var envParams EnvParameterValues

	// fmt.Println(">>",value)

	err := json.Unmarshal([]byte(value), &envParams)
	if err != nil {
		panic(err)
	}

	return envParams
}

func mustJsonToString(value EnvParameterValues) string {
	v, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	return string(v)
}
