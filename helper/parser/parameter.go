package parser

import (
	"errors"
	"reflect"

	"github.com/fatih/color"
)

func (parser *Parser) GetBackendParameterString(key string, optional bool) (error, string) {
	err, iface := parser.getSingleBackendParameterInterface(key, optional, "string")
	if err != nil {
		return err, ""
	}
	return nil, iface.(string)
}

func (parser *Parser) GetBackendParameterInt(key string, optional bool) (error, int) {
	err, iface := parser.getSingleBackendParameterInterface(key, optional, "int")
	if err != nil {
		return err, 0
	}
	return nil, iface.(int)
}

func (parser *Parser) GetBackendParameterBool(key string, optional bool) (error, bool) {
	err, iface := parser.getSingleBackendParameterInterface(key, optional, "bool")
	if err != nil {
		return err, false
	}
	return nil, iface.(bool)
}

func (parser *Parser) getSingleBackendParameterInterface(key string, optional bool, expectedType string) (error, interface{}) {
	switch parser.VarFileContent.(type) {
	case map[string]interface{}:
		mapped := parser.VarFileContent.(map[string]interface{})

		if mapped[key] == nil && optional == false {
			err := errors.New(key + " must be defined, was not found var-file")
			color.Red(err.Error())
			return err, nil
		}

		if reflect.TypeOf(mapped[key]).String() != expectedType {
			err := errors.New("Expected " + key + " to be " + expectedType + ", was " + reflect.TypeOf(mapped[key]).String())
			color.Red(err.Error())
			return err, nil
		}
		return nil, mapped[key]

	default:
		err := errors.New("Unknown var-file format")
		color.Red(err.Error())
		return err, nil
	}
}
