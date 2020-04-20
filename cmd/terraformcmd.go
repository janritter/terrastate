package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func getTerraformExecCmdForSubcommand(subcommand string, varFile string) *exec.Cmd {
	terraformExecutable, _ := exec.LookPath("terraform")

	cmd := &exec.Cmd{
		Path: terraformExecutable,
		Args: []string{
			terraformExecutable,
			subcommand,
			"--var-file=" + varFile},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
		Stdin:  os.Stdin,
	}

	fmt.Printf("\nRunning %s command: %s \n\n", subcommand, cmd.String())

	return cmd
}
