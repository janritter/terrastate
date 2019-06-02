package helper

import (
	"github.com/janritter/terrastate/helper/parser"
)

func GetBackendType(in interface{}) (string, error) {
	varParser := parser.NewParser(in)
	backend, _, err := varParser.GetBackendParameterString("state_backend", false)
	if err != nil {
		return "", err
	}
	return backend, nil
}
