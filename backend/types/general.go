package types

type stateFileAttributes struct {
}

type StateFileAttribute struct {
	AttributeKey string
	VarKey       string
	ExpectedType string
	Required     bool
	Value        interface{}
	Given        bool
}
