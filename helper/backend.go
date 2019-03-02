package helper

import (
	"errors"
	"log"
	"reflect"
)

func GetBackendType(in interface{}) (string, error) {
	var stateBucket string
	switch in.(type) {
	case map[string]interface{}:
		mapped := in.(map[string]interface{})
		if mapped["state_backend"] == nil {
			err := errors.New("state_backend must be defined")
			log.Println(err)
			return "", err
		}
		if reflect.TypeOf(mapped["state_backend"]).String() != "string" {
			err := errors.New("state_backend must be of type string, was " + reflect.TypeOf(mapped["state_backend"]).String())
			log.Println(err)
			return "", err
		}
		stateBucket = mapped["state_backend"].(string)

	default:
		err := errors.New("Unknown var-file format")
		log.Println(err)
		return "", err
	}

	return stateBucket, nil
}
