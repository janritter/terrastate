package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeVarFile(t *testing.T) {
	content := []byte(`
# comment
name  = "value with {{ current.dir }}"
flag  = true
off   = false
count = 3
ratio = 1.5
`)

	decoded, err := decodeVarFile(content, "test.tfvars")

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"name":  "value with {{ current.dir }}",
		"flag":  true,
		"off":   false,
		"count": 3,
		"ratio": 1.5,
	}, decoded)
}

func TestDecodeVarFileEmpty(t *testing.T) {
	decoded, err := decodeVarFile([]byte(""), "empty.tfvars")

	assert.NoError(t, err)
	assert.Empty(t, decoded)
}

func TestDecodeVarFileNull(t *testing.T) {
	decoded, err := decodeVarFile([]byte("key = null"), "null.tfvars")

	assert.NoError(t, err)
	assert.Nil(t, decoded["key"])
}

func TestDecodeVarFileSyntaxError(t *testing.T) {
	_, err := decodeVarFile([]byte(`key = "unterminated`), "bad.tfvars")

	assert.Error(t, err)
}

func TestDecodeVarFileCollections(t *testing.T) {
	content := []byte(`
azs  = ["a", "b"]
tags = { env = "test", count = 2 }
nested = {
  list = [1, true]
  obj  = { inner = "x" }
}
`)

	decoded, err := decodeVarFile(content, "collections.tfvars")

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"azs":  []interface{}{"a", "b"},
		"tags": map[string]interface{}{"env": "test", "count": 2},
		"nested": map[string]interface{}{
			"list": []interface{}{1, true},
			"obj":  map[string]interface{}{"inner": "x"},
		},
	}, decoded)
}

func TestDecodeVarFileUnknownValue(t *testing.T) {
	_, err := decodeVarFile([]byte(`key = var.foo`), "ref.tfvars")

	assert.Error(t, err)
}

func TestDecodeVarFileRepoTestFile(t *testing.T) {
	content, err := os.ReadFile("../test/test.tfvars")
	assert.NoError(t, err)

	decoded, err := decodeVarFile(content, "test.tfvars")

	assert.NoError(t, err)
	assert.Equal(t, "s3", decoded["state_backend"])
	assert.Equal(t, true, decoded["state_auto_remove_old"])
	assert.Equal(t, "terrastate/{{ current.dir }}/terraform.tfstate", decoded["state_key"])
}

// Regression test for https://github.com/janritter/terrastate/issues/93
func TestDecodeVarFileNumericMapKeys(t *testing.T) {
	content := []byte(`
aurora_instances = {
  1 = {
    instance_class = "db.t4g.medium"
  }
}
`)

	decoded, err := decodeVarFile(content, "numeric-keys.tfvars")

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"aurora_instances": map[string]interface{}{
			"1": map[string]interface{}{"instance_class": "db.t4g.medium"},
		},
	}, decoded)
}
