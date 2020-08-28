package parser

import "github.com/janritter/terrastate/backend/types"

type ParserAPI interface {
	Gather(stateFileAttributes []*types.StateFileAttribute)
	GatherTerrastateVariables(terrastateAttribute map[string]*types.TerrastateAttribute)
}

type Parser struct {
	VarFileContent interface{}
}

func NewParser(varFileContent interface{}) *Parser {
	return &Parser{
		VarFileContent: varFileContent,
	}
}
