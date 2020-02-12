package s3

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallParserForBackendParametersNoVarFileContent(t *testing.T) {
	stateConfigResult := stateConfig{}

	err := callParserForBackendParameters(nil, &stateConfigResult)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("Unknown var-file format"), err)
}

func TestCallParserForBackendParametersSuccess(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_bucket"] = "test_bucket"
	testMap["state_dynamodb_table"] = "test_dynamodb_table"
	testMap["state_key"] = "test/{{current.dir}}/terraform.tfstate"
	testMap["region"] = "eu-central-1"

	// Current dir gets replaced with the folder the test is located in, in this case "s3"
	expectedStateKey := "test/s3/terraform.tfstate"

	stateConfigResult := stateConfig{}

	err := callParserForBackendParameters(testMap, &stateConfigResult)
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, testMap["state_bucket"], stateConfigResult.Bucket)
	assert.Equal(t, testMap["state_dynamodb_table"], stateConfigResult.DynamoDBTable)
	assert.Equal(t, expectedStateKey, stateConfigResult.Key)
	assert.Equal(t, testMap["region"], stateConfigResult.Region)
}

func TestCallParserForBackendParametersMissingDynamoDB(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_bucket"] = "test_bucket"
	testMap["state_key"] = "test/{{current.dir}}/terraform.tfstate"
	testMap["region"] = "eu-central-1"

	stateConfigResult := stateConfig{}

	err := callParserForBackendParameters(testMap, &stateConfigResult)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("state_dynamodb_table must be defined, was not found in var-file"), err)
}

func TestCallParserForBackendParametersInvalidDynamoDB(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_bucket"] = "test_bucket"
	testMap["state_dynamodb_table"] = 0
	testMap["state_key"] = "test/{{current.dir}}/terraform.tfstate"
	testMap["region"] = "eu-central-1"

	stateConfigResult := stateConfig{}

	err := callParserForBackendParameters(testMap, &stateConfigResult)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("Expected state_dynamodb_table to be string, was int"), err)
}

func TestCallParserForBackendParametersMissingBucket(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_dynamodb_table"] = "test_bucket_table"
	testMap["state_key"] = "test/{{current.dir}}/terraform.tfstate"
	testMap["region"] = "eu-central-1"

	stateConfigResult := stateConfig{}

	err := callParserForBackendParameters(testMap, &stateConfigResult)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("state_bucket must be defined, was not found in var-file"), err)
}

func TestCallParserForBackendParametersInvalidBucket(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_bucket"] = 0
	testMap["state_dynamodb_table"] = "test_bucket_table"
	testMap["state_key"] = "test/{{current.dir}}/terraform.tfstate"
	testMap["region"] = "eu-central-1"

	stateConfigResult := stateConfig{}

	err := callParserForBackendParameters(testMap, &stateConfigResult)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("Expected state_bucket to be string, was int"), err)
}

func TestCallParserForBackendParametersMissingKey(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_bucket"] = "test_bucket"
	testMap["state_dynamodb_table"] = "test_bucket_table"
	testMap["region"] = "eu-central-1"

	stateConfigResult := stateConfig{}

	err := callParserForBackendParameters(testMap, &stateConfigResult)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("state_key must be defined, was not found in var-file"), err)
}

func TestCallParserForBackendParametersInvalidKey(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_bucket"] = "test_bucket"
	testMap["state_dynamodb_table"] = "test_bucket_table"
	testMap["state_key"] = 0
	testMap["region"] = "eu-central-1"

	stateConfigResult := stateConfig{}

	err := callParserForBackendParameters(testMap, &stateConfigResult)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("Expected state_key to be string, was int"), err)

	testMap = make(map[string]interface{})
	testMap["state_bucket"] = "test_bucket"
	testMap["state_dynamodb_table"] = "test_bucket_table"
	testMap["state_key"] = "test/terraform.tfstate"
	testMap["region"] = "eu-central-1"

	err = callParserForBackendParameters(testMap, &stateConfigResult)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("{{current.dir}} is missing the state_key parameter"), err)
}

func TestCallParserForBackendParametersMissingRegion(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_bucket"] = "test_bucket"
	testMap["state_dynamodb_table"] = "test_bucket_table"
	testMap["state_key"] = "test/{{current.dir}}/terraform.tfstate"

	stateConfigResult := stateConfig{}

	err := callParserForBackendParameters(testMap, &stateConfigResult)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("region must be defined, was not found in var-file"), err)
}

func TestCallParserForBackendParametersInvalidRegion(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_bucket"] = "test_bucket"
	testMap["state_dynamodb_table"] = "test_bucket_table"
	testMap["state_key"] = "test/{{current.dir}}/terraform.tfstate"
	testMap["region"] = 0

	stateConfigResult := stateConfig{}

	err := callParserForBackendParameters(testMap, &stateConfigResult)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("Expected region to be string, was int"), err)
}
