package dockerwrapper

import "strings"

type DockerWrapper interface {
	PullImage(image, tag string) error
	Run(runList []string, image, tag string) error
}

type wrapper struct {
	runner CommandRunner
}

func New(runner CommandRunner) DockerWrapper {
	return wrapper{runner: runner}
}

func (w wrapper) PullImage(image, tag string) error {
	_, err := w.runner.Run("pull " + image + ":" + tag)
	return err
}

func (w wrapper) Run(runList []string, image, tag string) error {
	dockerCommand := strings.Join(runList, " && ")
	command := "run --tty -i --rm " + image + ":" + tag + " '" + dockerCommand + "'"
	_, err := w.runner.Run(command)
	return err
}
