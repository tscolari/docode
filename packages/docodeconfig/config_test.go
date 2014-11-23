package docodeconfig_test

import (
	. "io/ioutil"

	"github.com/tscolari/docode/packages/docodeconfig"

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
		var subject docodeconfig.Configuration

		JustBeforeEach(func() {
			subject = docodeconfig.NewFromFile(sampleDocodeFile)
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

})
