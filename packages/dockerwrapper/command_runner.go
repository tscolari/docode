package dockerwrapper

import (
	"os"
	"os/exec"
)

type CommandRunner interface {
	Run(command string, args []string) error
}

type dockerRunner struct{}

func NewDockerRunner() dockerRunner {
	return dockerRunner{}
}

func (r dockerRunner) Run(command string, args []string) error {
	commandArgs := append([]string{command}, args...)
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
