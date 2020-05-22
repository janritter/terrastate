package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStringAfterSettingPlaceholderValuesSinglePlaceholder(t *testing.T) {
	result := GetStringAfterSettingPlaceholderValues("test/{{current.dir}}/test")

	assert.Equal(t, result, "test/helper/test")
}

func TestGetStringAfterSettingPlaceholderValuesMultiplePlaceholder(t *testing.T) {
	result := GetStringAfterSettingPlaceholderValues("test/{{current.dir}}/test/{{current.dir}}")

	assert.Equal(t, result, "test/helper/test/helper")
}

func TestGetCurrentDir(t *testing.T) {
	result := getCurrentDir()

	assert.Equal(t, result, "helper")
}
