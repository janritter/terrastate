package helper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBackendType(t *testing.T) {
	varFileContent := map[string]interface{}{
		"state_backend": "s3",
	}

	backendType, err := GetBackendType(varFileContent)

	assert.Equal(t, "s3", backendType)
	assert.Nil(t, err)
}

func TestGetBackendTypeMissingBackend(t *testing.T) {
	varFileContent := map[string]interface{}{}

	backendType, err := GetBackendType(varFileContent)

	assert.Equal(t, "", backendType)
	assert.Equal(t, errors.New("state_backend must be defined, was not found in var-file"), err)
}
