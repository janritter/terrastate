package parser

import "github.com/janritter/terrastate/backend/types"

type ParserAPI interface {
	Process([]*types.StateFileAttribute)
}

type Parser struct {
	VarFileContent interface{}
}

func NewParser(varFileContent interface{}) *Parser {
	return &Parser{
		VarFileContent: varFileContent,
	}
}
