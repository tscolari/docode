package main

import (
	"./docode"
	"./docodeconfig"

	"fmt"
)

func main() {
	config := docodeconfig.NewFromFile("./DocodeFile")
	runner := docode.New(config)
	err := runner.Run()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}
