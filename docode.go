package main

import (
	"github.com/tscolari/docode/packages/docode"
	"github.com/tscolari/docode/packages/docodeconfig"

	"flag"
)

func main() {
	docodeFilePath := flag.String("c", "./DocodeFile", "ConfigFile to load")
	argsConfig := fetchConfigFromArgs()

	flag.Parse()

	fileConfig := docodeconfig.NewFromFile(*docodeFilePath)
	config := docodeconfig.MergeConfigurations(argsConfig, fileConfig)
	runner := docode.New(config)
	err := runner.Run()
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}

func fetchConfigFromArgs() docodeconfig.ArgsConfiguration {
	argsConfig := docodeconfig.ArgsConfiguration{}
	argsConfig.SSHKey = flag.String("k", "", "Ssh key path to use")

	return argsConfig
}
