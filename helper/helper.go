package helper

import "github.com/janritter/terrastate/backend/types"

type HelperAPI interface {
	GetBackendType(in interface{}) (string, error)
	PrintStateFileAttributes(attributes []*types.StateFileAttribute)
	RemoveDotTerraformFolder(in interface{})
}

type Helper struct {
}

func NewHelper() *Helper {
	return &Helper{}
}
