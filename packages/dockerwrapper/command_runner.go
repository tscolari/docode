package dockerwrapper

import "os/exec"
import "strings"

type CommandRunner interface {
	Run(command string) (string, error)
}

type dockerRunner struct{}

func NewDockerRunner() dockerRunner {
	return dockerRunner{}
}

func (r dockerRunner) Run(command string) (string, error) {
	args := strings.Split(command, " ")
	cmd := exec.Command(r.prefix(), args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func (r *dockerRunner) prefix() string {
	return "docker"
}
