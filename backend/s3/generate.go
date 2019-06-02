package s3

import (
	"github.com/janritter/terrastate/helper"
)

func (backend *S3Backend) GenerateConfigurationForBackend(in interface{}) error {
	stateParams := stateConfig{}
	err := callParserForBackendParameters(in, &stateParams)
	if err != nil {
		return err
	}

	helper.PrintStateValues(stateParams)

	err = helper.RemoveDotTerraformFolder(in)
	if err != nil {
		return err
	}

	err = createBackendConfigurationFile(stateParams)
	return err
}
