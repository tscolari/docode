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
	return w.runner.Run("pull", []string{image + ":" + tag})
}

func (w wrapper) Run(runList []string, portMappings map[int]int, image, tag string) error {
	dockerCommand := strings.Join(runList, "&&")

	args := append(w.defaultStaticParams(), w.portsMapToArgsParams(portMappings)...)
	args = append(args, w.mountPointParams()...)
	args = append(args, image+":"+tag, "-c", dockerCommand)

	return w.runner.Run("run", args)
}

func (w wrapper) defaultStaticParams() []string {
	return []string{"--tty", "-i", "--rm", "-w", "/workdir", "--entrypoint", "/bin/sh"}
}

func (w wrapper) portsMapToArgsParams(portMappings map[int]int) []string {
	ports := []string{}

	for hostPort, dockerPort := range portMappings {
		ports = append(ports, []string{"-p", fmt.Sprintf("%d:%d", hostPort, dockerPort)}...)
	}

	return ports
}

func (w wrapper) mountPointParams() []string {
	workingDir, _ := filepath.Abs("")

	return []string{
		"-v",
		fmt.Sprintf("%s:/workdir", workingDir),
	}
}
