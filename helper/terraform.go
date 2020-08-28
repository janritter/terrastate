package helper

import (
	"os"

	"github.com/fatih/color"
)

func (helper *Helper) RemoveDotTerraformFolder(shouldRemove bool) {
	if shouldRemove {
		err := os.RemoveAll(".terraform/terraform.tfstate")
		if err != nil {
			color.Red(err.Error())
			osExit(1)
		}
		color.Green("Removed .terraform/terraform.tfstate")
		return
	}

	color.Blue("Skipping removing of .terraform/terraform.tfstate")
}
