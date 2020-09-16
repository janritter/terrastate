package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func getTerraformExecCmdForSubcommand(subcommand string, varFile string, terraformFlags string) *exec.Cmd {
	terraformExecutable, _ := exec.LookPath("terraform")

	args := []string{
		terraformExecutable,
		subcommand,
		"--var-file=" + varFile}

	if terraformFlags != "" {
		args = append(args, terraformFlags)
	}

	cmd := &exec.Cmd{
		Path:   terraformExecutable,
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
		Stdin:  os.Stdin,
	}

	fmt.Printf("\nRunning %s command: %s \n\n", subcommand, cmd.String())

	return cmd
}
