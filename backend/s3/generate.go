package s3

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fatih/color"
)

func (backend *S3Backend) GenerateStatefileForBackend(in interface{}) error {
	stateParams := stateConfig{}
	err := parseBackendParameter(in, &stateParams)
	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("------- Using the following values -------")
	fmt.Println("bucket = " + stateParams.Bucket)
	fmt.Println("dyanmodb_table = " + stateParams.DynamoDBTable)
	fmt.Println("key = " + stateParams.Key)
	fmt.Println("region = " + stateParams.Region)
	fmt.Println("------------------------------------------")

	err = createStateFile(stateParams)
	return err
}

func createStateFile(in stateConfig) error {
	fileContent :=
		`terraform {
  backend "s3" {
	encrypt        = true
	bucket         = "` + in.Bucket + `"
	region         = "` + in.Region + `"
	key            = "` + in.Key + `"
	dynamodb_table = "` + in.DynamoDBTable + `"
  }
}
`

	data := []byte(fileContent)
	err := ioutil.WriteFile("terrastate.tf", data, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	color.Green("Successfully created terrastate.tf")
	return nil
}
