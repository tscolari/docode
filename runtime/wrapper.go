package runtime

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Wrapper interface {
	PullImage(image, tag string) error
	Run(runList []string, portMappings map[int]int, image, tag, sshKey string, envSets, mountSets map[string]string) error
}

type wrapper struct {
	runner CommandRunner
}

func New() Wrapper {
	runner := runner{}
	return wrapper{runner: runner}
}

func NewWithRunner(commandRunner CommandRunner) Wrapper {
	return wrapper{runner: commandRunner}
}

func (w wrapper) PullImage(image, tag string) error {
	return w.runner.Run("pull", []string{image + ":" + tag})
}

func (w wrapper) Run(runList []string, portMappings map[int]int, image, tag, sshKey string, envSets, mountSets map[string]string) error {
	dockerCommand := strings.Join(runList, "&&")

	args := append(w.defaultStaticParams(), w.portsMapToArgsParams(portMappings)...)
	args = append(args, w.mountPointParams()...)
	args = append(args, w.environmentMappings(envSets)...)

	if len(sshKey) > 0 {
		args = append(args, w.mountSSHKey(sshKey)...)
		dockerCommand = "eval `ssh-agent -s`&&ssh-add /ssh_key&&" + dockerCommand
	}

	args = append(args, w.mountMappings(mountSets)...)
	args = append(args, image+":"+tag, "-c", dockerCommand)
	return w.runner.Run("run", args)
}

func (w wrapper) defaultStaticParams() []string {
	return []string{"--tty", "-i", "--rm", "-w", "/workdir", "--entrypoint", "/bin/sh"}
}

func (w wrapper) environmentMappings(envSets map[string]string) []string {
	mappings := []string{}

	for key, value := range envSets {
		mappings = append(mappings, "-e", key+"="+value)
	}

	return mappings
}

func (w wrapper) mountMappings(mountSets map[string]string) []string {
	mappings := []string{}

	for key, value := range mountSets {
		mappings = append(mappings, "-v", key+":"+value)
	}

	return mappings
}

func (w wrapper) mountSSHKey(sshKeyPath string) []string {
	return []string{
		"-v",
		sshKeyPath + ":/ssh_key",
	}
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
