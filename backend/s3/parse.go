package s3

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func parseBackendParameter(in interface{}, out *stateConfig) error {
	switch in.(type) {
	case map[string]interface{}:
		mapped := in.(map[string]interface{})

		if mapped["state_bucket"] == nil {
			err := errors.New("state_bucket must be defined")
			log.Println(err)
			return err
		}
		if reflect.TypeOf(mapped["state_bucket"]).String() != "string" {
			err := errors.New("state_bucket must be of type string, was " + reflect.TypeOf(mapped["state_bucket"]).String())
			log.Println(err)
			return err
		}
		out.Bucket = mapped["state_bucket"].(string)

		if mapped["state_dynamodb_table"] == nil {
			err := errors.New("state_dynamodb_table must be defined")
			log.Println(err)
			return err
		}
		if reflect.TypeOf(mapped["state_dynamodb_table"]).String() != "string" {
			err := errors.New("state_dynamodb_table must be of type string, was " + reflect.TypeOf(mapped["state_dynamodb_table"]).String())
			log.Println(err)
			return err
		}
		out.DynamoDBTable = mapped["state_dynamodb_table"].(string)

		if mapped["state_key"] == nil {
			err := errors.New("state_key must be defined")
			log.Println(err)
			return err
		}
		if reflect.TypeOf(mapped["state_key"]).String() != "string" {
			err := errors.New("state_key must be of type string, was " + reflect.TypeOf(mapped["state_key"]).String())
			log.Println(err)
			return err
		}

		// Remove all spaces
		key := strings.ReplaceAll(mapped["state_key"].(string), " ", "")

		// Check if the key string contains current.dir
		if !strings.Contains(key, "{{current.dir}}") {
			err := errors.New("{{current.dir}} is missing the state_key string")
			log.Println(err)
			return err
		}

		// Replace placeholder with current dir
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		dir := filepath.Base(path)
		fmt.Println("Current Directory = " + dir)
		key = strings.ReplaceAll(key, "{{current.dir}}", dir)

		out.Key = key

		if mapped["region"] == nil {
			err := errors.New("region must be defined")
			log.Println(err)
			return err
		}
		if reflect.TypeOf(mapped["region"]).String() != "string" {
			err := errors.New("region must be of type string, was " + reflect.TypeOf(mapped["region"]).String())
			log.Println(err)
			return err
		}
		out.Region = mapped["region"].(string)

	default:
		err := errors.New("Unknown var-file format")
		log.Println(err)
		return err
	}

	return nil
}
