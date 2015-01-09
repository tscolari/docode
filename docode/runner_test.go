package docode_test

import (
	"github.com/tscolari/docode/config"
	. "github.com/tscolari/docode/docode"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeDockerWrapper struct {
	image        string
	tag          string
	pulled       bool
	runList      []string
	portMappings map[int]int
	envSets      map[string]string
	mountSets    map[string]string
}

func (w *fakeDockerWrapper) PullImage(image, tag string) error {
	w.image = image
	w.tag = tag
	w.pulled = true
	return nil
}

func (w *fakeDockerWrapper) Run(runList []string, portMappings map[int]int, image, tag, sshKey string, envSets, mountSets map[string]string) error {
	w.image = image
	w.tag = tag
	w.runList = runList
	w.portMappings = portMappings
	w.envSets = envSets
	w.mountSets = mountSets
	return nil
}

var _ = Describe("runner", func() {
	var wrapper *fakeDockerWrapper
	var runner *Runner
	var configuration config.Configuration

	Describe(".Run", func() {
		JustBeforeEach(func() {
			wrapper = &fakeDockerWrapper{}
			configuration = config.Configuration{
				ImageName: "busybox",
				ImageTag:  "oldone",
				DontPull:  true,
				RunList:   []string{"ls", "cd tmp"},
				Ports:     map[int]int{2222: 1111},
				EnvSets:   map[string]string{"HELLO": "world"},
				MountSets: map[string]string{"/home": "/other_home", "/tmp": "/trash"},
			}

			runner = NewWithWrapper(configuration, wrapper)
		})

		It("sets correctly the docker image", func() {
			runner.Run()
			Expect(wrapper.image).To(Equal("busybox"))
		})

		It("sets correctly the docker tag", func() {
			runner.Run()
			Expect(wrapper.tag).To(Equal("oldone"))
		})

		It("sets correctly the runList", func() {
			runner.Run()
			Expect(wrapper.runList).To(Equal([]string{"ls", "cd tmp"}))
		})

		It("sets correctly the port mappings", func() {
			runner.Run()
			Expect(wrapper.portMappings).To(Equal(map[int]int{2222: 1111}))
		})

		It("sets correctly the envs", func() {
			runner.Run()
			Expect(wrapper.envSets).To(Equal(map[string]string{"HELLO": "world"}))
		})

		It("doesn't pull the image if dont_pull is true", func() {
			runner.Run()
			Expect(wrapper.pulled).To(Equal(false))
		})

		It("sets correctly the mountings", func() {
			runner.Run()
			Expect(wrapper.mountSets).To(Equal(map[string]string{"/home": "/other_home", "/tmp": "/trash"}))
		})

		Context("when dont_pull is not set or not present", func() {
			It("pulls the image if dont_pull false", func() {
				configuration.DontPull = false
				runner = NewWithWrapper(configuration, wrapper)
				runner.Run()
				Expect(wrapper.pulled).To(Equal(true))
			})
		})
	})
})
