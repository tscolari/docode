package docodeconfig

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type configuration struct {
	ImageName string      `yaml:"image_name"`
	ImageTag  string      `yaml:"image_tag"`
	Ports     map[int]int `yaml:"ports"`
	RunList   []string    `yaml:"run_list"`
}

func NewFromFile(filename string) configuration {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Failed to open configuration file: " + filename)
	}

	config := configuration{}
	config.loadString(contents)
	return config
}

func (c *configuration) loadString(configContents []byte) {
	yaml.Unmarshal(configContents, &c)
}
