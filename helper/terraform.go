package helper

import (
	"os"

	"github.com/fatih/color"
	"github.com/janritter/terrastate/helper/parser"
)

func RemoveDotTerraformFolder(in interface{}) {
	varParser := parser.NewParser(in)
	shouldRemove, valueSet, err := varParser.GetBackendParameterBool("state_auto_remove_old", true)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	if valueSet == false || shouldRemove == false {
		color.Blue("Skipping removing of .terraform/terraform.tfstate")
	}

	if valueSet == true && shouldRemove == true {
		err = os.RemoveAll(".terraform/terraform.tfstate")
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		color.Green("Removed .terraform/terraform.tfstate")
	}
}
