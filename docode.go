package main

import (
	"github.com/tscolari/docode/config"
	"github.com/tscolari/docode/docode"

	"flag"
)

func main() {
	docodeFilePath := flag.String("c", "./DocodeFile", "ConfigFile to load")
	argsConfig := fetchConfigFromArgs()

	flag.Parse()

	fileConfig := config.NewFromFile(*docodeFilePath)
	configuration := config.MergeConfigurations(argsConfig, fileConfig)
	runner := docode.New(configuration)
	err := runner.Run()
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}

func fetchConfigFromArgs() config.ArgsConfiguration {
	argsConfig := config.ArgsConfiguration{}

	argsConfig.SSHKey = flag.String("k", "", "Ssh key path to use")
	argsConfig.ImageName = flag.String("i", "", "Image name to use")
	argsConfig.ImageTag = flag.String("t", "", "Image tag to use")
	argsConfig.DontPull = flag.Bool("n", false, "Skip pulling the image")

	return argsConfig
}
