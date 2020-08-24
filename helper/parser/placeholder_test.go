package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStringAfterSettingPlaceholderValuesSinglePlaceholder(t *testing.T) {
	result := GetStringAfterSettingPlaceholderValues("test/{{current.dir}}/test")

	assert.Equal(t, "test/helper/test", result)
}

func TestGetStringAfterSettingPlaceholderValuesMultiplePlaceholder(t *testing.T) {
	result := GetStringAfterSettingPlaceholderValues("test/{{current.dir}}/test/{{current.dir}}")

	assert.Equal(t, "test/helper/test/helper", result)
}

func TestGetCurrentDir(t *testing.T) {
	result := getCurrentDir()

	assert.Equal(t, "helper", result)
}
