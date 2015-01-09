package runtime

import (
	"os"
	"os/exec"
)

type CommandRunner interface {
	Run(command string, args []string) error
}

type runner struct{}

func NewDockerRunner() runner {
	return runner{}
}

func (r runner) Run(command string, args []string) error {
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

func (r *runner) prefix() string {
	return "docker"
}
