package dockerwrapper_test

import (
	"../dockerwrapper"
	"errors"
	"fmt"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeCommandRunner struct {
	shouldError     bool
	receivedCommand string
}

func NewFakerCommandRunner() *fakeCommandRunner {
	return &fakeCommandRunner{shouldError: false}
}

func NewFailingCommandRunner() fakeCommandRunner {
	return fakeCommandRunner{shouldError: true}
}

func (r *fakeCommandRunner) Run(command string) error {
	if r.shouldError {
		return errors.New("Failed to run")
	}

	r.receivedCommand = command
	return nil
}

var _ = Describe("Wrapper", func() {
	var wrapper dockerwrapper.DockerWrapper
	var commandRunner *fakeCommandRunner

	Describe(".PullImage", func() {

		JustBeforeEach(func() {
			commandRunner = NewFakerCommandRunner()
			wrapper = dockerwrapper.New(commandRunner)
		})

		It("sends the correct parameters to command runner", func() {
			wrapper.PullImage("busybox", "latest")
			Ω(commandRunner.receivedCommand).To(Equal("pull busybox:latest"))
		})
	})

	Describe(".Run", func() {
		It("sends the correct parameters to command runner", func() {
			wrapper.Run([]string{"bundle install", "tmux"}, map[int]int{22: 2022, 80: 8080}, "busybox", "latest")
			workingFolder, _ := filepath.Abs("")
			expectedCommand := fmt.Sprintf("run --tty -i --rm -w /workdir --entrypoint /bin/sh -p 22:2022 -p 80:8080 -v %s:/workdir busybox:latest -c |bundle install&&tmux", workingFolder)
			Ω(commandRunner.receivedCommand).To(Equal(expectedCommand))
		})
	})
})
