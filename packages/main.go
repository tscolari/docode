package main

import (
	"./docodeconfig"

	"fmt"
)

func main() {
	config := docodeconfig.NewConfigurationFromFile("./DocodeFile")
	fmt.Printf("--- config:\n%v\n\n", config)
}
