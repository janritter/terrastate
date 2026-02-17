package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStringAfterSettingPlaceholderValuesSinglePlaceholder(t *testing.T) {
	result := getStringAfterSettingPlaceholderValues("test/{{current.dir}}/test")

	assert.Equal(t, "test/parser/test", result)
}

func TestGetStringAfterSettingPlaceholderValuesMultiplePlaceholder(t *testing.T) {
	result := getStringAfterSettingPlaceholderValues("test/{{current.dir}}/test/{{current.dir}}")

	assert.Equal(t, "test/parser/test/parser", result)
}

func TestGetStringAfterSettingPlaceholderValuesParentDir(t *testing.T) {
	result := getStringAfterSettingPlaceholderValues("test/{{parent.dir}}/test")

	assert.Equal(t, "test/helper/test", result)
}

func TestGetCurrentDir(t *testing.T) {
	result := getCurrentDir()

	assert.Equal(t, "parser", result)
}

func TestGetParentDir(t *testing.T) {
	result := getParentDir()

	assert.Equal(t, "helper", result)
}

func TestParentDirPlaceholderWithTestFile(t *testing.T) {
	// Read the test-parent.tfvars file
	testFilePath := filepath.Join("..", "..", "test", "test-parent.tfvars")
	content, err := os.ReadFile(testFilePath)
	assert.NoError(t, err)

	// Process the content to replace placeholders
	processedContent := getStringAfterSettingPlaceholderValues(string(content))

	// Check that parent.dir placeholders were replaced correctly
	assert.Contains(t, processedContent, "helper-terraform-state")
	assert.Contains(t, processedContent, "helper-lock-table")
	assert.Contains(t, processedContent, "helper/terraform.tfstate")
}

func TestCombinedPlaceholdersWithTestFile(t *testing.T) {
	// Read the test-parent-current-combined.tfvars file
	testFilePath := filepath.Join("..", "..", "test", "test-parent-current-combined.tfvars")
	content, err := os.ReadFile(testFilePath)
	assert.NoError(t, err)

	// Process the content to replace placeholders
	processedContent := getStringAfterSettingPlaceholderValues(string(content))

	// Check that both parent.dir and current.dir placeholders were replaced correctly
	assert.Contains(t, processedContent, "terrastate/helper/parser/terraform.tfstate")
}
