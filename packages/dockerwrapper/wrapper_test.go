package dockerwrapper_test

import (
	"../dockerwrapper"
	"fmt"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeCommandRunner struct {
	receivedCommand string
	receivedArgs    []string
}

func NewFakerCommandRunner() *fakeCommandRunner {
	return &fakeCommandRunner{}
}

func (r *fakeCommandRunner) Run(command string, args []string) error {
	r.receivedCommand = command
	r.receivedArgs = args
	return nil
}

var _ = Describe("Wrapper", func() {
	var wrapper dockerwrapper.DockerWrapper
	var commandRunner *fakeCommandRunner

	JustBeforeEach(func() {
		commandRunner = NewFakerCommandRunner()
		wrapper = dockerwrapper.NewWithRunner(commandRunner)
	})

	Describe(".PullImage", func() {
		It("sends the correct parameters to command runner", func() {
			wrapper.PullImage("busybox", "latest")
			Ω(commandRunner.receivedCommand).To(Equal("pull"))
			Ω(commandRunner.receivedArgs).To(Equal([]string{"busybox:latest"}))
		})
	})

	Describe(".Run", func() {
		Context("When all arguments are supplied", func() {
			It("sends the correct parameters to command runner", func() {
				wrapper.Run([]string{"bundle install", "tmux"}, map[int]int{22: 2022, 80: 8080}, "busybox", "latest", "my_ssh_key")
				Ω(commandRunner.receivedCommand).To(Equal("run"))

				workingFolder, _ := filepath.Abs("")
				expectedArgs := []string{
					"--tty",
					"-i",
					"--rm",
					"-w",
					"/workdir",
					"--entrypoint",
					"/bin/sh",
					"-p",
					"22:2022",
					"-p",
					"80:8080",
					"-v",
					fmt.Sprintf("%s:/workdir", workingFolder),
					"-v",
					"my_ssh_key:/ssh_key",
					"busybox:latest",
					"-c",
					"eval `ssh-agent -s`&&ssh-add /ssh_key&&bundle install&&tmux",
				}
				Ω(commandRunner.receivedArgs).To(Equal(expectedArgs))
			})
		})

		Context("When no ssh-key is given", func() {
			It("should not mount anything to ssh_key nor add ssh-add on it", func() {
				wrapper.Run([]string{"tmux"}, map[int]int{}, "busybox", "latest", "")
				Ω(commandRunner.receivedCommand).To(Equal("run"))

				workingFolder, _ := filepath.Abs("")
				expectedArgs := []string{
					"--tty",
					"-i",
					"--rm",
					"-w",
					"/workdir",
					"--entrypoint",
					"/bin/sh",
					"-v",
					fmt.Sprintf("%s:/workdir", workingFolder),
					"busybox:latest",
					"-c",
					"tmux",
				}
				Ω(commandRunner.receivedArgs).To(Equal(expectedArgs))
			})
		})
	})
})
