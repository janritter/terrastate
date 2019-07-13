package s3

import (
	"os"
	"text/template"

	"github.com/fatih/color"
)

func createBackendConfigurationFile(in stateConfig) error {
	t, err := template.New("backend").Parse(`terraform {
  backend "s3" {
	encrypt        = true
	{{ if .Bucket }}bucket         = "{{ .Bucket }}"{{ end }}
	{{ if .Region }}region         = "{{ .Region }}"{{ end }}
	{{ if .Key }}key            = "{{ .Key }}"{{ end }}
	{{ if .DynamoDBTable }}dynamodb_table = "{{ .DynamoDBTable }}"{{ end }}
  }
}
	  `)
	if err != nil {
		color.Red(err.Error())
		return err
	}

	f, err := os.Create("terrastate.tf")
	if err != nil {
		color.Red(err.Error())
		return err
	}

	err = t.Execute(f, in)
	if err != nil {
		color.Red(err.Error())
		return err
	}

	color.Green("Successfully created terrastate.tf")
	return nil
}
