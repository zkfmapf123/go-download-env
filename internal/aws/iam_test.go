package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUserRole(t *testing.T) {
	
	awsParams := MustNewAWS()
	
	role, err :=awsParams.GetUser()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, role, "admin")
}

func Test_AWSConfig(t *testing.T) {
	awsParams := MustNewAWS()
	assert.Equal(t, awsParams.role, "admin")
}
