package main

import (
	"github.com/tscolari/docode/packages/docode"
	"github.com/tscolari/docode/packages/docodeconfig"

	"flag"
)

func main() {
	docodeFilePath := flag.String("config", "./DocodeFile", "ConfigFile to load")
	flag.Parse()

	fileConfig := docodeconfig.NewFromFile(*docodeFilePath)
	runner := docode.New(fileConfig)
	err := runner.Run()
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
