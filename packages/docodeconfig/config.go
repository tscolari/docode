package docodeconfig

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	ImageName string            `yaml:"image_name"`
	ImageTag  string            `yaml:"image_tag"`
	Ports     map[int]int       `yaml:"ports"`
	RunList   []string          `yaml:"run_list"`
	SSHKey    string            `yaml:"ssh_key"`
	DontPull  bool              `yaml:"dont_pull"`
	EnvSets   map[string]string `yaml:"env"`
	MountSets map[string]string `yaml:"mount"`
}

func NewFromFile(filename string) Configuration {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Failed to open configuration file: %s", filename)
	}

	config := Configuration{}
	config.loadString(contents)
	return config
}

func MergeConfigurations(mainConfig ArgsConfiguration, secondaryConfig Configuration) Configuration {
	config := Configuration{}

	if mainConfig.ImageName == nil || *mainConfig.ImageName == "" {
		config.ImageName = secondaryConfig.ImageName
	} else {
		config.ImageName = *mainConfig.ImageName
	}

	if mainConfig.ImageTag == nil || *mainConfig.ImageTag == "" {
		config.ImageTag = secondaryConfig.ImageTag
	} else {
		config.ImageTag = *mainConfig.ImageTag
	}

	if mainConfig.SSHKey == nil || *mainConfig.SSHKey == "" {
		config.SSHKey = secondaryConfig.SSHKey
	} else {
		config.SSHKey = *mainConfig.SSHKey
	}

	if mainConfig.Ports == nil {
		config.Ports = secondaryConfig.Ports
	} else {
		config.Ports = *mainConfig.Ports
	}

	if mainConfig.RunList == nil {
		config.RunList = secondaryConfig.RunList
	} else {
		config.RunList = *mainConfig.RunList
	}

	if mainConfig.EnvSets == nil {
		config.EnvSets = secondaryConfig.EnvSets
	} else {
		config.EnvSets = *mainConfig.EnvSets
	}

	if mainConfig.MountSets == nil {
		config.MountSets = secondaryConfig.MountSets
	} else {
		config.MountSets = *mainConfig.MountSets
	}

	if mainConfig.DontPull == nil || !*mainConfig.DontPull {
		config.DontPull = secondaryConfig.DontPull
	} else {
		config.DontPull = *mainConfig.DontPull
	}

	return config
}

func (c *Configuration) loadString(configContents []byte) {
	yaml.Unmarshal(configContents, &c)
}
