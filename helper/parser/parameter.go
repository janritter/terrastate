package parser

import (
	"errors"
	"os"
	"reflect"

	"github.com/fatih/color"
	"github.com/janritter/terrastate/backend/types"
)

func (parser *Parser) Process(stateFileAttributes []*types.StateFileAttribute) {
	errorRun := false
	switch parser.VarFileContent.(type) {
	case map[string]interface{}:
		mapped := parser.VarFileContent.(map[string]interface{})

		for _, attribute := range stateFileAttributes {
			attribute.Given = mapped[attribute.VarKey] != nil

			if attribute.Required && !attribute.Given {
				err := errors.New(attribute.VarKey + " must be defined, was not found in the var-file")
				color.Red(err.Error())
				errorRun = true
			}

			if attribute.Given {
				if reflect.TypeOf(mapped[attribute.VarKey]).String() != attribute.ExpectedType {
					err := errors.New("Expected " + attribute.VarKey + " to be " + attribute.ExpectedType + ", was " + reflect.TypeOf(mapped[attribute.VarKey]).String())
					color.Red(err.Error())
					errorRun = true
				} else {
					attribute.Value = mapped[attribute.VarKey]
				}
			}
		}

		if errorRun {
			os.Exit(1)
		}

	default:
		err := errors.New("Unknown var-file format")
		color.Red(err.Error())
		os.Exit(1)
	}
}

func (parser *Parser) GetBackendParameterString(key string, optional bool) (string, bool, error) {
	iface, valueSet, err := parser.getSingleBackendParameterInterface(key, optional, "string")
	if err != nil {
		return "", valueSet, err
	}
	if valueSet == false {
		return "", valueSet, nil
	}
	return iface.(string), valueSet, nil
}

func (parser *Parser) GetBackendParameterInt(key string, optional bool) (int, bool, error) {
	iface, valueSet, err := parser.getSingleBackendParameterInterface(key, optional, "int")
	if err != nil {
		return 0, valueSet, err
	}
	if valueSet == false {
		return 0, valueSet, nil
	}
	return iface.(int), valueSet, nil
}

// Return values are: an error, the parsed bool value and if the value in the var-file was set.
func (parser *Parser) GetBackendParameterBool(key string, optional bool) (bool, bool, error) {
	iface, valueSet, err := parser.getSingleBackendParameterInterface(key, optional, "bool")
	if err != nil {
		return false, valueSet, err
	}
	if valueSet == false {
		return false, valueSet, nil
	}
	return iface.(bool), valueSet, nil
}

func (parser *Parser) getSingleBackendParameterInterface(key string, optional bool, expectedType string) (interface{}, bool, error) {
	switch parser.VarFileContent.(type) {
	case map[string]interface{}:
		mapped := parser.VarFileContent.(map[string]interface{})

		valueSet := mapped[key] != nil

		if valueSet == false {
			if optional == false {
				err := errors.New(key + " must be defined, was not found in var-file")
				color.Red(err.Error())
				return nil, valueSet, err
			}
			return nil, valueSet, nil
		}

		if reflect.TypeOf(mapped[key]).String() != expectedType {
			err := errors.New("Expected " + key + " to be " + expectedType + ", was " + reflect.TypeOf(mapped[key]).String())
			color.Red(err.Error())
			return nil, false, err
		}
		return mapped[key], valueSet, nil

	default:
		err := errors.New("Unknown var-file format")
		color.Red(err.Error())
		return nil, false, err
	}
}
