package docode_test

import (
	"../docode"
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
run_list:
  -
    tmux
`

	sampleDocodeFile =
		writeTemporaryDocodeFile(yamlContents)

	Describe("NewConfigurationFromFile", func() {

		It("correctly maps the `image_name`", func() {
			subject := docode.NewConfigurationFromFile(sampleDocodeFile)
			Expect(subject.ImageName).To(Equal("docode-base"))
		})

		It("correctly maps the `image_tag`", func() {
			subject := docode.NewConfigurationFromFile(sampleDocodeFile)
			Expect(subject.ImageTag).To(Equal("latest"))
		})

		It("correctly maps `ports`", func() {
			subject := docode.NewConfigurationFromFile(sampleDocodeFile)
			Expect(subject.Ports).To(Equal(map[int]int{80: 80}))
		})

		It("correctly maps `run_list`", func() {
			subject := docode.NewConfigurationFromFile(sampleDocodeFile)
			Expect(subject.RunList).To(Equal([]string{"tmux"}))
		})
	})

})
