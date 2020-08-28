package types

type StateFileAttribute struct {
	AttributeKey string
	VarKey       string
	ExpectedType string
	Required     bool
	Value        interface{}
	Given        bool
}
