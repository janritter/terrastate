package s3

import (
	"github.com/janritter/terrastate/backend/types"
	helperAPI "github.com/janritter/terrastate/helper"
	creatorAPI "github.com/janritter/terrastate/helper/creator"
	parserAPI "github.com/janritter/terrastate/helper/parser"
)

var backendAttributes = types.BackendAttributes{
	StateFileAttributes: []*types.StateFileAttribute{
		{
			AttributeKey: "bucket",
			VarKey:       "state_bucket",
			ExpectedType: "string",
			Required:     true,
		},
		{
			AttributeKey: "key",
			VarKey:       "state_key",
			ExpectedType: "string",
			Required:     true,
		},
		{
			AttributeKey: "region",
			VarKey:       "region",
			ExpectedType: "string",
			Required:     true,
		},
	},
}

type S3BackendAPI interface {
	Generate() error
}

type S3Backend struct {
	parser  parserAPI.ParserAPI
	creator creatorAPI.CreatorAPI
	helper  helperAPI.HelperAPI
}

func NewS3Backend(parser parserAPI.ParserAPI, creator creatorAPI.CreatorAPI, helper helperAPI.HelperAPI, terrastateAttributes map[string]*types.TerrastateAttribute) *S3Backend {
	backendAttributes.TerrastateAttributes = terrastateAttributes

	return &S3Backend{
		parser:  parser,
		creator: creator,
		helper:  helper,
	}
}

func (backend *S3Backend) Generate() {
	backend.parser.Gather(backendAttributes.StateFileAttributes)

	backend.helper.PrintStateFileAttributes(backendAttributes.StateFileAttributes)

	backend.helper.RemoveDotTerraformFolder(backendAttributes.TerrastateAttributes["state_auto_remove_old"].Value.(bool))

	backend.creator.Create(backendAttributes.StateFileAttributes, "s3")
}
