package helper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func ReplacePlaceholderInStateKey(stateKey string) (error, string) {
	// Remove all spaces
	key := strings.ReplaceAll(stateKey, " ", "")

	// Check if the key string contains current.dir
	if !strings.Contains(key, "{{current.dir}}") {
		err := errors.New("{{current.dir}} is missing the state_key parameter")
		color.Red(err.Error())
		return err, ""
	}

	// Replace placeholder with current dir
	path, err := os.Getwd()
	if err != nil {
		color.Red(err.Error())
		return err, ""
	}
	dir := filepath.Base(path)
	fmt.Println("Current Directory = " + dir)
	key = strings.ReplaceAll(key, "{{current.dir}}", dir)
	return nil, key
}
