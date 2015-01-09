package config_test

import (
	. "io/ioutil"

	"github.com/tscolari/docode/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	var writeTemporaryDocodeFile = func(content string) string {
		file, err := TempFile("", "DocodeFile")
		if err != nil {
			panic("Error trying to create temporary file")
		}

		_, err = file.WriteString(content)
		if err != nil {
			panic("Writing temporary DocodeFile")
		}

		return file.Name()
	}

	var sampleDocodeFile string

	yamlContents := `
image_name: docode-base
image_tag: latest
dont_pull: true
ssh_key: /some_key
ports:
  80: 80
  22: 1022
run_list:
  - memcached -d
  - tmux
`

	sampleDocodeFile =
		writeTemporaryDocodeFile(yamlContents)

	Describe(".NewFromFile", func() {
		var subject config.Configuration

		JustBeforeEach(func() {
			subject = config.NewFromFile(sampleDocodeFile)
		})

		It("correctly maps the `image_name`", func() {
			Expect(subject.ImageName).To(Equal("docode-base"))
		})

		It("correctly maps the `image_tag`", func() {
			Expect(subject.ImageTag).To(Equal("latest"))
		})

		It("correctly maps `ports`", func() {
			Expect(subject.Ports).To(Equal(map[int]int{80: 80, 22: 1022}))
		})

		It("correctly maps `run_list`", func() {
			Expect(subject.RunList).To(Equal([]string{"memcached -d", "tmux"}))
		})

		It("correctly maps `ssh_key`", func() {
			Expect(subject.SSHKey).To(Equal("/some_key"))
		})

		It("correctly maps `dont_pull`", func() {
			Expect(subject.DontPull).To(Equal(true))
		})
	})

	Describe("MergeConfigurations", func() {
		var configA config.ArgsConfiguration
		var configB config.Configuration
		var mergedConfig config.Configuration

		BeforeEach(func() {
			imageName := "imageA"
			SSHKey := "/id_rsa"
			imageTag := ""
			configA = config.ArgsConfiguration{
				ImageName: &imageName,
				ImageTag:  &imageTag,
				SSHKey:    &SSHKey,
				Ports: &map[int]int{
					80: 80,
				},
			}

			configB = config.Configuration{
				ImageName: "imageB",
				ImageTag:  "tagB",
				Ports:     map[int]int{80: 90, 100: 100},
				EnvSets:   map[string]string{"HOME": "/me"},
				MountSets: map[string]string{"/tmp": "/tmpb"},
				DontPull:  true,
			}
		})

		JustBeforeEach(func() {
			mergedConfig = config.MergeConfigurations(configA, configB)
		})

		Context("precedence", func() {
			It("gives preference to configA", func() {
				Expect(mergedConfig.ImageName).To(Equal("imageA"))
				Expect(mergedConfig.SSHKey).To(Equal("/id_rsa"))
				Expect(mergedConfig.Ports).To(Equal(map[int]int{80: 80}))
			})
		})

		Context("Missing keys", func() {
			It("uses configB values", func() {
				Expect(mergedConfig.ImageTag).To(Equal("tagB"))
				Expect(mergedConfig.EnvSets).To(Equal(map[string]string{"HOME": "/me"}))
				Expect(mergedConfig.MountSets).To(Equal(map[string]string{"/tmp": "/tmpb"}))
				Expect(mergedConfig.DontPull).To(Equal(true))
			})
		})
	})
})
