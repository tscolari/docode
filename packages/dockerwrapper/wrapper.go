package dockerwrapper

import (
	"fmt"
	"path/filepath"
	"strings"
)

type DockerWrapper interface {
	PullImage(image, tag string) error
	Run(runList []string, portMappings map[int]int, image, tag string) error
}

type wrapper struct {
	runner CommandRunner
}

func New(runner CommandRunner) DockerWrapper {
	return wrapper{runner: runner}
}

func (w wrapper) PullImage(image, tag string) error {
	return w.runner.Run("pull " + image + ":" + tag)
}

func (w wrapper) Run(runList []string, portMappings map[int]int, image, tag string) error {
	dockerCommand := strings.Join(runList, "&&")
	dockerParams := "--tty -i --rm -w /workdir --entrypoint /bin/sh"
	ports := ""

	for hostPort, dockerPort := range portMappings {
		ports = fmt.Sprintf("%s -p %d:%d", ports, hostPort, dockerPort)
	}

	workingDir, _ := filepath.Abs("")
	mountPoint := fmt.Sprintf("-v %s:/workdir", workingDir)

	command := fmt.Sprintf("run %s%s %s %s:%s -c |%s", dockerParams, ports, mountPoint, image, tag, dockerCommand)
	return w.runner.Run(command)
}
