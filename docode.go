package main

import (
	"github.com/tscolari/docode/packages/docode"
	"github.com/tscolari/docode/packages/docodeconfig"

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
