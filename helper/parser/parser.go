package parser

type ParserAPI interface {
	GetBackendParameterString(key string, optional bool) (string, bool, error)
	GetBackendParameterInt(key string, optional bool) (int, bool, error)
	GetBackendParameterBool(key string, optional bool) (bool, bool, error)
}

type Parser struct {
	VarFileContent interface{}
}

func NewParser(varFileContent interface{}) *Parser {
	return &Parser{
		VarFileContent: varFileContent,
	}
}
