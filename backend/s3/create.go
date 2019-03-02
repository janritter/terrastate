package s3

import (
	"io/ioutil"
	"log"

	"github.com/fatih/color"
)

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
