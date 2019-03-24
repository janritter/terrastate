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
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/janritter/terrastate/backend"
	"github.com/janritter/terrastate/helper"

	"github.com/janritter/terrastate/backend/iface"

	"github.com/hashicorp/hcl"
	"github.com/spf13/cobra"
)

type backends struct {
	iface.BackendAPI
}

var varFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terrastate",
	Short: "Tool to manage multiple state backends in Terraform - Allows Multi account setups",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if varFile == "" {
			log.Fatal("--var-file Flag must be set")
		}

		fmt.Println("Using var-file = " + varFile)

		b, err := ioutil.ReadFile(varFile)
		if err != nil {
			log.Fatal(err)
		}

		var decoded interface{}
		err = hcl.Decode(&decoded, string(b))
		if err != nil {
			log.Fatal(err)
		}

		backendType, err := helper.GetBackendType(decoded)
		if err != nil {
			os.Exit(1)
		}

		backendInterface, err := backend.GetBackendInterface(backendType)
		if err != nil {
			os.Exit(1)
		}

		backendBase := backends{
			BackendAPI: backendInterface,
		}

		backendBase.GenerateConfigurationForBackend(decoded)
	},
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
