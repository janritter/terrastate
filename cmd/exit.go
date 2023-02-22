package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// passThroughExitCode checks if the error from the executed terraform command starts with "exit status" and exits terrastate with the same exit code as the terraform command
func passThroughExitCode(err error) {
	// Check if error starts with "exit status" and pass through the exit code
	if strings.HasPrefix(err.Error(), "exit status") {
		exitCode, convertErr := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		if convertErr != nil {
			color.Red("Error converting non zero exit code from terraform command to int: " + convertErr.Error() + " - exiting with exit code 1")
			os.Exit(1)
		}
		os.Exit(exitCode)
	}
}
