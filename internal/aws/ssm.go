package aws

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const (
	KEY = "/goenvs/settings"
)

type EnvParameterValues struct {
	Teams []string `json:"teams"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (a *AWSEnvParmas) GetParameter() (EnvParameterValues, error) {

	resp, err := a.ssmParameterClient.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: aws.String(KEY),
	})

	if err != nil {

		// Create
		if strings.Contains(err.Error(), "ParameterNotFound") {
	
			if err = a.createParameter(); err != nil {
				return EnvParameterValues{} ,err
			}
	
			return EnvParameterValues{}, nil
		}
		
		return EnvParameterValues{}, err
	}

	return mustStirngToJson(*resp.Parameter.Value), nil
}

func (a *AWSEnvParmas) createParameter() error {
	
	envValues := EnvParameterValues{
		Teams: []string{},
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	_, err := a.ssmParameterClient.PutParameter(context.TODO(), &ssm.PutParameterInput{
		Name: aws.String(KEY),
		Value: aws.String(mustJsonToString(envValues)),
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
