package helper

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestFolderAndFile() string {
	os.Mkdir(".terraform", os.ModePerm)
	path, _ := filepath.Abs("./.terraform/terraform.tfstate")
	os.Create(path)

	return path
}

func removeTestFolderAndFile() {
	path, _ := filepath.Abs("./.terraform")
	os.RemoveAll(path)
}

func TestRemoveDotTerraformFolderDeactivated(t *testing.T) {
	varFileContent := map[string]interface{}{
		"state_auto_remove_old": false,
	}
	path := createTestFolderAndFile()

	err := RemoveDotTerraformFolder(varFileContent)

	assert.Nil(t, err)
	assert.FileExists(t, path)

	removeTestFolderAndFile()
}

func TestRemoveDotTerraformFolderValueNotSet(t *testing.T) {
	varFileContent := map[string]interface{}{
		"state_auto_remove_old": false,
	}
	path := createTestFolderAndFile()

	err := RemoveDotTerraformFolder(varFileContent)

	assert.Nil(t, err)
	assert.FileExists(t, path)

	removeTestFolderAndFile()
}

func TestRemoveDotTerraformFolder(t *testing.T) {
	varFileContent := map[string]interface{}{
		"state_auto_remove_old": true,
	}
	path := createTestFolderAndFile()

	err := RemoveDotTerraformFolder(varFileContent)

	assert.Nil(t, err)
	assert.NoFileExists(t, path)

	removeTestFolderAndFile()
}
