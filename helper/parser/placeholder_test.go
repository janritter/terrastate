package parser

import (
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

func TestGetCurrentDir(t *testing.T) {
	result := getCurrentDir()

	assert.Equal(t, "parser", result)
}
