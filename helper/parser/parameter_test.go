package parser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// GetBackendParameterString
//
func TestGetBackendParameterString(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["testKey"] = "testValue"

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterString("testKey", false)

	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, "testValue", result)
	assert.Equal(t, true, valueSet)
}

func TestGetBackendParameterStringWrongType(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["testKey"] = 0

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterString("testKey", false)

	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("Expected testKey to be string, was int"), err)
	assert.Equal(t, "", result)
	assert.Equal(t, false, valueSet)
}

func TestGetBackendParameterStringNotSetOptional(t *testing.T) {
	testMap := make(map[string]interface{})

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterString("testKey", true)

	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, "", result)
	assert.Equal(t, false, valueSet)
}

func TestGetBackendParameterStringNotSetNotOptional(t *testing.T) {
	testMap := make(map[string]interface{})

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterString("testKey", false)

	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("testKey must be defined, was not found var-file"), err)
	assert.Equal(t, "", result)
	assert.Equal(t, false, valueSet)
}

//
// GetBackendParameterInt
//

func TestGetBackendParameterInt(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["testKey"] = 10

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterInt("testKey", false)

	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, 10, result)
	assert.Equal(t, true, valueSet)
}

func TestGetBackendParameterIntWrongType(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["testKey"] = "10"

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterInt("testKey", false)

	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("Expected testKey to be int, was string"), err)
	assert.Equal(t, 0, result)
	assert.Equal(t, false, valueSet)
}

func TestGetBackendParameterIntNotSetOptional(t *testing.T) {
	testMap := make(map[string]interface{})

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterInt("testKey", true)

	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, 0, result)
	assert.Equal(t, false, valueSet)
}

func TestGetBackendParameterIntNotSetNotOptional(t *testing.T) {
	testMap := make(map[string]interface{})

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterInt("testKey", false)

	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("testKey must be defined, was not found var-file"), err)
	assert.Equal(t, 0, result)
	assert.Equal(t, false, valueSet)
}

//
// GetBackendParameterBool
//

func TestGetBackendParameterBool(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["testKey"] = false

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterBool("testKey", false)

	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, false, result)
	assert.Equal(t, true, valueSet)
}

func TestGetBackendParameterBoolWrongType(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["testKey"] = "false"

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterBool("testKey", false)

	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("Expected testKey to be bool, was string"), err)
	assert.Equal(t, false, result)
	assert.Equal(t, false, valueSet)
}

func TestGetBackendParameterBoolNotSetOptional(t *testing.T) {
	testMap := make(map[string]interface{})

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterBool("testKey", true)

	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, false, result)
	assert.Equal(t, false, valueSet)
}

func TestGetBackendParameterBoolNotSetNotOptional(t *testing.T) {
	testMap := make(map[string]interface{})

	parser := NewParser(testMap)
	result, valueSet, err := parser.GetBackendParameterBool("testKey", false)

	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("testKey must be defined, was not found var-file"), err)
	assert.Equal(t, false, result)
	assert.Equal(t, false, valueSet)
}

//
// getSingleBackendParameterInterface
//

func TestGetSingleBackendParameterInterfaceInvalidFileContent(t *testing.T) {
	parser := NewParser(nil)
	result, valueSet, err := parser.getSingleBackendParameterInterface("testKey", false, "string")

	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("Unknown var-file format"), err)
	assert.Equal(t, nil, result)
	assert.Equal(t, false, valueSet)
}
