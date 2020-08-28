package backend

import (
	"errors"
	"log"

	"github.com/janritter/terrastate/backend/iface"
	"github.com/janritter/terrastate/backend/s3"
	"github.com/janritter/terrastate/backend/types"
	helperAPI "github.com/janritter/terrastate/helper"
	creatorAPI "github.com/janritter/terrastate/helper/creator"
	parserAPI "github.com/janritter/terrastate/helper/parser"
)

var terrastateAttributes = map[string]*types.TerrastateAttribute{
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

func GetBackendInterface(varFile interface{}) (iface.BackendAPI, error) {
	parser := parserAPI.NewParser(varFile)
	creator := creatorAPI.NewCreator()
	helper := helperAPI.NewHelper()

	parser.GatherTerrastateVariables(terrastateAttributes)

	backendType := terrastateAttributes["state_backend"].Value.(string)

	switch backendType {
	case "s3":
		return s3.NewS3Backend(parser, creator, helper, terrastateAttributes), nil
	}

	err := errors.New("Backend " + backendType + " is currently not supported")
	log.Println(err)
	return nil, err
}
