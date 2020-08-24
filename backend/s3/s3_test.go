package s3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewS3Backend(t *testing.T) {
	backend := NewS3Backend(nil)

	assert.Empty(t, backend, "Expected to be empty")
	assert.NotNil(t, backend, "Expected not to be nil")
}
