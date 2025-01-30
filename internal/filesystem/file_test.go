package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getCurrenta(t *testing.T){
	dir := GetCurrentPath()
	
	assert.NotNil(t, dir)
}
