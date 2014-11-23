package docode_test

import (
	. "github.com/tscolari/docode/packages/docode"
	"github.com/tscolari/docode/packages/docodeconfig"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeDockerWrapper struct {
	image        string
	tag          string
	runList      []string
	portMappings map[int]int
}

func (w *fakeDockerWrapper) PullImage(image, tag string) error {
	w.image = image
	w.tag = tag
	return nil
}

func (w *fakeDockerWrapper) Run(runList []string, portMappings map[int]int, image, tag, sshKey string) error {
	w.image = image
	w.tag = tag
	w.runList = runList
	w.portMappings = portMappings
	return nil
}

var _ = Describe("runner", func() {
	var wrapper *fakeDockerWrapper
	var runner *Runner
	var config docodeconfig.Configuration

	Describe(".Run", func() {
		JustBeforeEach(func() {
			wrapper = &fakeDockerWrapper{}
			config = docodeconfig.Configuration{
				ImageName: "busybox",
				ImageTag:  "oldone",
				RunList:   []string{"ls", "cd tmp"},
				Ports:     map[int]int{2222: 1111},
			}

			runner = NewWithWrapper(config, wrapper)
		})

		It("sends commands to the docker wrapper", func() {
			runner.Run()
			Expect(wrapper.image).To(Equal("busybox"))
			Expect(wrapper.tag).To(Equal("oldone"))
			Expect(wrapper.runList).To(Equal([]string{"ls", "cd tmp"}))
			Expect(wrapper.portMappings).To(Equal(map[int]int{2222: 1111}))
		})
	})
})
