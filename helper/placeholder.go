package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type placeholder struct {
	key   string
	value string
}

var placeholders = []placeholder{
	{
		key:   "current.dir",
		value: getCurrentDir(),
	},
}

func getCurrentDir() string {
	path, err := os.Getwd()
	if err != nil {
		color.Red(err.Error())
	}
	return filepath.Base(path)
}

func GetStringAfterSettingPlaceholderValues(input string) string {
	output := strings.ReplaceAll(input, " ", "")

	for _, placeholder := range placeholders {
		if strings.Contains(output, "{{"+placeholder.key+"}}") {
			fmt.Println("PLACEHOLDER - replacing '{{" + placeholder.key + "}}' in '" + output + "' with '" + placeholder.value + "'")
		}
		output = strings.ReplaceAll(output, "{{"+placeholder.key+"}}", placeholder.value)
	}

	return output
}
