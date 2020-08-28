package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHelper(t *testing.T) {
	helper := NewHelper()
	assert.NotNil(t, helper)
}
