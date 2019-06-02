package s3

import (
	"github.com/janritter/terrastate/helper"

	"github.com/janritter/terrastate/helper/parser"
)

func callParserForBackendParameters(in interface{}, out *stateConfig) error {
	varParser := parser.NewParser(in)

	bucket, _, err := varParser.GetBackendParameterString("state_bucket", false)
	if err != nil {
		return err
	}
	out.Bucket = bucket

	dynamodbTable, _, err := varParser.GetBackendParameterString("state_dynamodb_table", false)
	if err != nil {
		return err
	}
	out.DynamoDBTable = dynamodbTable

	stateKey, _, err := varParser.GetBackendParameterString("state_key", false)
	if err != nil {
		return err
	}
	err, stateKey = helper.ReplacePlaceholderInStateKey(stateKey)
	if err != nil {
		return err
	}
	out.Key = stateKey

	region, _, err := varParser.GetBackendParameterString("region", false)
	if err != nil {
		return err
	}
	out.Region = region

	return nil
}
