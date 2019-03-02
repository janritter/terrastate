package s3

import (
	"github.com/janritter/terrastate/helper"
)

func (backend *S3Backend) GenerateStatefileForBackend(in interface{}) error {
	stateParams := stateConfig{}
	err := parseBackendParameter(in, &stateParams)
	if err != nil {
		return err
	}

	helper.PrintStateValues(stateParams)

	err = createStateFile(stateParams)
	return err
}
