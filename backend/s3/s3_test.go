package s3

import (
	"testing"

	"github.com/janritter/terrastate/backend/types"
	"github.com/janritter/terrastate/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var terrastateTestAttributes = map[string]*types.TerrastateAttribute{
	"state_auto_remove_old": {
		ExpectedType: "bool",
		Required:     false,
		Value:        true,
	},
}

func TestNewS3Backend(t *testing.T) {
	parser := new(mocks.ParserAPI)
	helper := new(mocks.HelperAPI)
	creator := new(mocks.CreatorAPI)

	backend := NewS3Backend(parser, creator, helper, terrastateTestAttributes)

	assert.NotEmpty(t, backend, "Expected to be not empty")
	assert.NotNil(t, backend, "Expected not to be nil")
	assert.Equal(t, parser, backend.parser)
	assert.Equal(t, helper, backend.helper)
	assert.Equal(t, creator, backend.creator)
}

func TestGenerateSuccess(t *testing.T) {
	parser := new(mocks.ParserAPI)
	parser.On("Gather", mock.AnythingOfType("[]*types.StateFileAttribute")).Return(nil, nil)

	helper := new(mocks.HelperAPI)
	helper.On("PrintStateFileAttributes", mock.AnythingOfType("[]*types.StateFileAttribute")).Return(nil, nil)
	helper.On("RemoveDotTerraformFolder", mock.AnythingOfType("bool")).Return(nil, nil)

	creator := new(mocks.CreatorAPI)
	creator.On("Create", mock.AnythingOfType("[]*types.StateFileAttribute"), mock.AnythingOfType("string")).Return(nil, nil)

	backend := NewS3Backend(parser, creator, helper, terrastateTestAttributes)

	backend.Generate()

	parser.AssertExpectations(t)
	helper.AssertExpectations(t)
	creator.AssertExpectations(t)
}
