package dockerwrapper_test

import (
	"../dockerwrapper"
	"errors"

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

func (r *fakeCommandRunner) Run(command string) (string, error) {
	if r.shouldError {
		return "", errors.New("Failed to run")
	}

	r.receivedCommand = command
	return "Ran", nil
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
			wrapper.Run([]string{"bundle install", "tmux"}, "busybox", "latest")
			Ω(commandRunner.receivedCommand).To(Equal("run --tty -i --rm busybox:latest 'bundle install && tmux'"))
		})
	})
})
