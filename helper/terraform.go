package helper

import (
	"os"

	"github.com/fatih/color"
	"github.com/janritter/terrastate/helper/parser"
)

func RemoveDotTerraformFolder(in interface{}) error {
	varParser := parser.NewParser(in)
	shouldRemove, valueSet, err := varParser.GetBackendParameterBool("state_auto_remove_old", true)
	if err != nil {
		return err
	}

	if valueSet == false {
		color.Blue("Skipping removing of .terraform folder")
		return nil
	}

	if valueSet == true && shouldRemove == true {
		err = os.RemoveAll(".terraform")
		if err != nil {
			color.Red(err.Error())
			return err
		}
		color.Green("Removed .terraform folder")
	}
	return nil
}
