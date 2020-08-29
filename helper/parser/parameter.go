package parser

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/fatih/color"
	"github.com/janritter/terrastate/backend/types"
)

func (parser *Parser) Gather(stateFileAttributes []*types.StateFileAttribute) {
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
					attribute.Given = false
					color.Red(err.Error())
					errorRun = true
				} else if reflect.TypeOf(mapped[attribute.VarKey]).String() == "string" {
					attribute.Value = getStringAfterSettingPlaceholderValues(fmt.Sprintf("%v", mapped[attribute.VarKey]))
				} else {
					attribute.Value = mapped[attribute.VarKey]
				}
			}
		}

		if errorRun {
			osExit(1)
		}

	default:
		err := errors.New("Unknown var-file format")
		color.Red(err.Error())
		osExit(1)
	}
}

func (parser *Parser) GatherTerrastateVariables(terrastateAttribute map[string]*types.TerrastateAttribute) {
	errorRun := false
	switch parser.VarFileContent.(type) {
	case map[string]interface{}:
		mapped := parser.VarFileContent.(map[string]interface{})

		for varKey, attribute := range terrastateAttribute {
			if attribute.Required && mapped[varKey] == nil {
				err := errors.New(varKey + " must be defined, was not found in the var-file")
				color.Red(err.Error())
				errorRun = true
			}

			if mapped[varKey] != nil {
				if reflect.TypeOf(mapped[varKey]).String() != attribute.ExpectedType {
					err := errors.New("Expected " + varKey + " to be " + attribute.ExpectedType + ", was " + reflect.TypeOf(mapped[varKey]).String())
					color.Red(err.Error())
					errorRun = true
				} else if reflect.TypeOf(mapped[varKey]).String() == "string" {
					attribute.Value = getStringAfterSettingPlaceholderValues(fmt.Sprintf("%v", mapped[varKey]))
				} else {
					attribute.Value = mapped[varKey]
				}
			}
		}

		if errorRun {
			osExit(1)
		}

	default:
		err := errors.New("Unknown var-file format")
		color.Red(err.Error())
		osExit(1)
	}
}
