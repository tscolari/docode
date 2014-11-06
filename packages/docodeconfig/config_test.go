package docodeconfig_test

import (
	"../docodeconfig"
	. "io/ioutil"

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

		It("correctly maps the `image_name`", func() {
			subject := docodeconfig.NewFromFile(sampleDocodeFile)
			Expect(subject.ImageName).To(Equal("docode-base"))
		})

		It("correctly maps the `image_tag`", func() {
			subject := docodeconfig.NewFromFile(sampleDocodeFile)
			Expect(subject.ImageTag).To(Equal("latest"))
		})

		It("correctly maps `ports`", func() {
			subject := docodeconfig.NewFromFile(sampleDocodeFile)
			Expect(subject.Ports).To(Equal(map[int]int{80: 80, 22: 1022}))
		})

		It("correctly maps `run_list`", func() {
			subject := docodeconfig.NewFromFile(sampleDocodeFile)
			Expect(subject.RunList).To(Equal([]string{"memcached -d", "tmux"}))
		})
	})

})
