package main

import (
	"./dockerwrapper"
	"./docodeconfig"

	"fmt"
)

func main() {
	config := docodeconfig.NewFromFile("./DocodeFile")
	dockerRunner := dockerwrapper.NewDockerRunner()
	wrapper := dockerwrapper.New(dockerRunner)
	fmt.Printf("--- Pulling image: %s:%s\n", config.ImageName, config.ImageTag)
	err := wrapper.PullImage(config.ImageName, config.ImageTag)
	if err != nil {
		fmt.Printf("	  ERROR: %s\n", err.Error())
	}
	fmt.Printf("--- Running container\n")
	err = wrapper.Run(config.RunList, config.Ports, config.ImageName, config.ImageTag)
	if err != nil {
		fmt.Printf("	  ERROR: %s\n", err.Error())
	}
}
