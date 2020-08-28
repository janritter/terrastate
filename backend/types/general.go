package types

type BackendAttributes struct {
	TerrastateAttributes map[string]*TerrastateAttribute
	StateFileAttributes  []*StateFileAttribute
}

type TerrastateAttribute struct {
	ExpectedType string
	Required     bool
	Value        interface{}
}

type StateFileAttribute struct {
	AttributeKey string
	VarKey       string
	ExpectedType string
	Required     bool
	Value        interface{}
	Given        bool
}
