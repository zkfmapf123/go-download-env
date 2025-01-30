package aws

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

const (
	TAG_KEY = "role"
	ADMIN = "admin"
	DEVELOPER = "developer"
)

func (a *AWSEnvParmas) GetUser() (string ,error) {
	
	arn, err := a.getCurrentUserUseContext()
	if err != nil {
		return "", err
	}

	username := strings.Split(arn, "/")[1]

	user, err := a.iamClient.GetUser(context.TODO(), &iam.GetUserInput{
		UserName: &username,
	})
	if err != nil {
		return "" ,err
	}

	for _ ,v := range user.User.Tags{
		
		if *v.Key == TAG_KEY {
			return *v.Value, nil
		}
	}
	
	return DEVELOPER, nil
}

func (a *AWSEnvParmas) getCurrentUserUseContext() (string, error) {
	
	result, err := a.stsClient.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		return "", fmt.Errorf("failed to get caller identity: %w", err)
	}

	return *result.Arn, nil
}