package helper

import "github.com/janritter/terrastate/backend/types"

type HelperAPI interface {
	PrintStateFileAttributes(attributes []*types.StateFileAttribute)
	RemoveDotTerraformFolder(shouldRemove bool)
}

type Helper struct {
}

func NewHelper() *Helper {
	return &Helper{}
}
