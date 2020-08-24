package s3

import (
	"github.com/janritter/terrastate/backend/types"
	"github.com/janritter/terrastate/helper"
	"github.com/janritter/terrastate/helper/creator"
	"github.com/janritter/terrastate/helper/parser"
)

var stateFileAttributes = []*types.StateFileAttribute{
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
}

type S3BackendAPI interface {
	Generate() error
}

type S3Backend struct {
	VarParser *parser.Parser
	VarFile   interface{}
}

func NewS3Backend(varFile interface{}) *S3Backend {
	return &S3Backend{
		VarParser: parser.NewParser(varFile),
		VarFile:   varFile,
	}
}

func (backend *S3Backend) Generate() {
	backend.VarParser.Process(stateFileAttributes)

	helper.PrintStateFileAttributes(stateFileAttributes)

	helper.RemoveDotTerraformFolder(backend.VarFile)

	fileCreator := creator.NewCreator()
	fileCreator.Create(stateFileAttributes, "s3")
}
