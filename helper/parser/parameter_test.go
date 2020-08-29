package parser

import (
	"testing"

	"github.com/janritter/terrastate/backend/types"
	"github.com/stretchr/testify/assert"
)

func TestGatherSuccess(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_bucket"] = "test_bucket"
	testMap["test_number"] = 0

	stateFileAttributes := []*types.StateFileAttribute{
		{
			AttributeKey: "bucket",
			VarKey:       "state_bucket",
			ExpectedType: "string",
			Required:     true,
		},
		{
			AttributeKey: "number",
			VarKey:       "test_number",
			ExpectedType: "int",
			Required:     true,
		},
		{
			AttributeKey: "bool",
			VarKey:       "bool_test",
			ExpectedType: "bool",
			Required:     false,
		},
	}

	parser := NewParser(testMap)

	parser.Gather(stateFileAttributes)

	assert.Equal(t, testMap["state_bucket"], stateFileAttributes[0].Value)
	assert.True(t, stateFileAttributes[0].Given)

	assert.Equal(t, testMap["test_number"], stateFileAttributes[1].Value)
	assert.True(t, stateFileAttributes[1].Given)
}

func TestGatherMissingRequiredParameter(t *testing.T) {
	testMap := make(map[string]interface{})

	stateFileAttributes := []*types.StateFileAttribute{
		{
			AttributeKey: "bucket",
			VarKey:       "state_bucket",
			ExpectedType: "string",
			Required:     true,
		},
	}

	// Check for osExit
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	testExit := func(code int) {
		got = code
	}
	osExit = testExit
	// End check for osExit

	parser := NewParser(testMap)

	parser.Gather(stateFileAttributes)

	assert.Nil(t, stateFileAttributes[0].Value)
	assert.False(t, stateFileAttributes[0].Given)
	assert.Equal(t, got, 1)
}

func TestGatherWrongParameterType(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["test_number"] = "not_a_number"

	stateFileAttributes := []*types.StateFileAttribute{
		{
			AttributeKey: "bucket",
			VarKey:       "test_number",
			ExpectedType: "int",
			Required:     true,
		},
	}

	// Check for osExit
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	testExit := func(code int) {
		got = code
	}
	osExit = testExit
	// End check for osExit

	parser := NewParser(testMap)

	parser.Gather(stateFileAttributes)

	assert.Nil(t, stateFileAttributes[0].Value)
	assert.False(t, stateFileAttributes[0].Given)
	assert.Equal(t, got, 1)
}

func TestGatherUnknownFormat(t *testing.T) {
	testMap := make(map[int]interface{})

	stateFileAttributes := []*types.StateFileAttribute{}

	// Check for osExit
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	testExit := func(code int) {
		got = code
	}
	osExit = testExit
	// End check for osExit

	parser := NewParser(testMap)

	parser.Gather(stateFileAttributes)

	assert.Equal(t, got, 1)
}

func TestGatherTerrastateVariablesSuccess(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_backend"] = "backend"

	terrastateAttributes := map[string]*types.TerrastateAttribute{
		"state_backend": {
			ExpectedType: "string",
			Required:     true,
		},
		"state_auto_remove_old": {
			ExpectedType: "bool",
			Required:     false,
			Value:        false, // default
		},
	}

	parser := NewParser(testMap)

	parser.GatherTerrastateVariables(terrastateAttributes)

	assert.Equal(t, testMap["state_backend"], terrastateAttributes["state_backend"].Value)
	assert.False(t, terrastateAttributes["state_auto_remove_old"].Value.(bool))
}

func TestGatherTerrastateVariablesMissingRequiredParameter(t *testing.T) {
	testMap := make(map[string]interface{})

	terrastateAttributes := map[string]*types.TerrastateAttribute{
		"state_backend": {
			ExpectedType: "string",
			Required:     true,
		},
	}

	// Check for osExit
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	testExit := func(code int) {
		got = code
	}
	osExit = testExit
	// End check for osExit

	parser := NewParser(testMap)

	parser.GatherTerrastateVariables(terrastateAttributes)

	assert.Equal(t, got, 1)
}

func TestGatherTerrastateVariablesWrongParameterType(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["state_backend"] = 0

	terrastateAttributes := map[string]*types.TerrastateAttribute{
		"state_backend": {
			ExpectedType: "string",
			Required:     true,
		},
	}

	// Check for osExit
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	testExit := func(code int) {
		got = code
	}
	osExit = testExit
	// End check for osExit

	parser := NewParser(testMap)

	parser.GatherTerrastateVariables(terrastateAttributes)

	assert.Nil(t, terrastateAttributes["state_backend"].Value)
	assert.Equal(t, got, 1)
}

func TestGatherTerrastateVariablesUnknownFormat(t *testing.T) {
	testMap := make(map[int]interface{})

	terrastateAttributes := map[string]*types.TerrastateAttribute{}

	// Check for osExit
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	testExit := func(code int) {
		got = code
	}
	osExit = testExit
	// End check for osExit

	parser := NewParser(testMap)

	parser.GatherTerrastateVariables(terrastateAttributes)

	assert.Equal(t, got, 1)
}
