package aws

import (
	"testing"
)

// func Test_GetParameter(t *testing.T) {

// 	awsConfig := MustNewAWS()

// 	v, err := awsConfig.GetParameter()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(v)
// }

func Test_CreateSecreManager(t *testing.T) {
	awsConfig := MustNewAWS()

	err := awsConfig.CreateSecretManager("test-a")
	if err != nil {
		panic(err)
	}
}

func Test_GetSecreManager(t *testing.T) {
	awsConfig := MustNewAWS()

	err := awsConfig.GetSecretManager("test-bbb")
	if err != nil {
		panic(err)
	}
}

