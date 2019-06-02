package parser

type ParserAPI interface {
	GetBackendParameterString(key string, optional bool) (error, string)
	GetBackendParameterInt(key string, optional bool) (error, int)
	GetBackendParameterBool(key string, optional bool) (error, bool)
}

type Parser struct {
	VarFileContent interface{}
}

func NewParser(varFileContent interface{}) *Parser {
	return &Parser{
		VarFileContent: varFileContent,
	}
}
