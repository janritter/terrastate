package s3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewS3Backend(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_bucket"] = "test_bucket"
	testMap["state_dynamodb_table"] = "test_dynamodb_table"

	backend := NewS3Backend(testMap)

	assert.NotEmpty(t, backend, "Expected to be not empty")
	assert.NotNil(t, backend, "Expected not to be nil")
	assert.NotNil(t, backend.VarParser, "Expected parser not to be nil")
	assert.Equal(t, testMap, backend.VarFile)
}
