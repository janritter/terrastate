/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Run terraform destroy through terrastate, also executes terrastate and terraform init",
	Run: func(cmd *cobra.Command, args []string) {
		// Terrastate
		fmt.Printf("Running terrastate \n\n")
		rootCmd.Run(cmd, args)

		// Terraform init
		if err := getTerraformExecCmdForSubcommand("init", varFile, "").Run(); err != nil {
			color.Red("terraform init returned the following error code: " + err.Error())
			return
		}

		// Terraform destroy
		if err := getTerraformExecCmdForSubcommand("destroy", varFile, strings.Join(args, " ")).Run(); err != nil {
			color.Red("terraform destroy returned the following error code: " + err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
