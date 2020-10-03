// Copyright © 2020 Jan Ritter <git@janrtr.de>
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
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Run terraform refresh through terrastate, also executes terrastate and terraform init",
	Run: func(cmd *cobra.Command, args []string) {
		// Terrastate
		fmt.Printf("Running terrastate \n\n")
		rootCmd.Run(cmd, args)

		// Terraform init
		if err := getTerraformExecCmdForSubcommand("init", varFile, "").Run(); err != nil {
			color.Red("terraform init returned the following error code: " + err.Error())
			return
		}

		// Terraform refresh
		if err := getTerraformExecCmdForSubcommand("refresh", varFile, strings.Join(args, " ")).Run(); err != nil {
			color.Red("terraform refresh returned the following error code: " + err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}
