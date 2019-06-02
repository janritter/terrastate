package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["testKey"] = "testValue"

	parser := NewParser(testMap)

	assert.Equal(t, testMap, parser.VarFileContent)
}
