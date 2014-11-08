package dockerwrapper

import (
	"os"
	"os/exec"
	"strings"
)

type CommandRunner interface {
	Run(command string) error
}

type dockerRunner struct{}

func NewDockerRunner() dockerRunner {
	return dockerRunner{}
}

func (r dockerRunner) Run(command string) error {
	args := strings.Split(command, "|")
	commandArgs := strings.Split(args[0], " ")
	if len(args) == 2 {
		commandArgs[len(commandArgs)-1] = commandArgs[len(commandArgs)-1] + " " + args[1]
	}

	cmd := exec.Command(r.prefix(), commandArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()

	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (r *dockerRunner) prefix() string {
	return "docker"
}
