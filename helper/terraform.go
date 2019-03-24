package helper

import (
	"errors"
	"log"
	"os"
	"reflect"

	"github.com/fatih/color"
)

func RemoveDotTerraformFolder(in interface{}) error {
	switch in.(type) {
	case map[string]interface{}:
		mapped := in.(map[string]interface{})
		if mapped["state_auto_remove_old"] == nil {
			color.Blue("Skipping removing of .terraform folder")

			return nil
		}
		if reflect.TypeOf(mapped["state_auto_remove_old"]).String() != "bool" {
			err := errors.New("state_auto_remove_old must be of type bool, was " + reflect.TypeOf(mapped["state_auto_remove_old"]).String())
			log.Println(err)

			return err
		}
		if !mapped["state_auto_remove_old"].(bool) {
			color.Blue("Skipping removing of .terraform folder")

			return nil
		}

	default:
		err := errors.New("Unknown var-file format")
		log.Println(err)

		return err
	}

	err := os.RemoveAll(".terraform")
	if err != nil {
		color.Red(err.Error())

		return err
	}

	color.Green("Removed .terraform folder")

	return nil
}
