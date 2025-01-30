package aws

import (
	"fmt"
	"testing"
)


func Test_GetParameter(t *testing.T) {

	awsConfig := MustNewAWS()
	
	v, err := awsConfig.GetParameter()
	if err != nil {
		panic(err)
	}

	fmt.Println(v)
}
