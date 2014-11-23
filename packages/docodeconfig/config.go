package docodeconfig

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	ImageName string      `yaml:"image_name"`
	ImageTag  string      `yaml:"image_tag"`
	Ports     map[int]int `yaml:"ports"`
	RunList   []string    `yaml:"run_list"`
	SSHKey    string      `yaml:"ssh_key"`
}

func NewFromFile(filename string) Configuration {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Failed to open configuration file: " + filename)
	}

	config := Configuration{}
	config.loadString(contents)
	return config
}

func (c *Configuration) loadString(configContents []byte) {
	yaml.Unmarshal(configContents, &c)
}
