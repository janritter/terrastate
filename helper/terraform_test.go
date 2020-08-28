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

func TestRemoveDotTerraformFolderFalset(t *testing.T) {
	path := createTestFolderAndFile()

	helper := NewHelper()
	helper.RemoveDotTerraformFolder(false)

	assert.FileExists(t, path)

	removeTestFolderAndFile()
}

func TestRemoveDotTerraformFolderTrue(t *testing.T) {
	path := createTestFolderAndFile()

	helper := NewHelper()
	helper.RemoveDotTerraformFolder(true)

	assert.NoFileExists(t, path)

	removeTestFolderAndFile()
}
