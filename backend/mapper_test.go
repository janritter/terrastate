package backend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBackendInterfaceSuccess(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_backend"] = "s3"

	backend, err := GetBackendInterface(testMap)

	assert.NoError(t, err)
	assert.NotNil(t, backend)
}

func TestGetBackendInterfaceUnsupportedBackend(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_backend"] = "not_supported"

	backend, err := GetBackendInterface(testMap)

	assert.Error(t, err)
	assert.Nil(t, backend)
}
