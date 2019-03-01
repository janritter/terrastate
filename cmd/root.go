// Copyright © 2019 Jan Ritter <git@janrtr.de>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/spf13/cobra"

	"github.com/fatih/color"
)

type state struct {
	Bucket        string
	DynamoDBTable string
	Key           string
	Region        string
}

var varFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terrastate",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if varFile == "" {
			log.Fatal("--var-file Flag must be set")
		}

		fmt.Println("Using var-file = " + varFile)

		b, err := ioutil.ReadFile(varFile) // just pass the file name
		if err != nil {
			log.Fatal(err)
		}

		var out interface{}

		err = hcl.Decode(&out, string(b))
		if err != nil {
			log.Fatal(err)
		}

		stateValues := state{}
		err = parseStateValues(out, &stateValues)
		if err != nil {
			os.Exit(1)
		}

		fmt.Println("")
		fmt.Println("------- Using the following values -------")
		fmt.Println("bucket = " + stateValues.Bucket)
		fmt.Println("dyanmodb_table = " + stateValues.DynamoDBTable)
		fmt.Println("key = " + stateValues.Key)
		fmt.Println("region = " + stateValues.Region)
		fmt.Println("------------------------------------------")

		err = generateTerraformStateFile(stateValues)
		if err != nil {
			os.Exit(1)
		}
	},
}

func parseStateValues(in interface{}, out *state) error {
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
	}

	return nil
}

func generateTerraformStateFile(in state) error {
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&varFile, "var-file", "", "Terraform variables file")
}
