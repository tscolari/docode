package main

import (
	"./docode"

	"fmt"
)

func main() {
	config := docode.NewConfigurationFromFile("./DocodeFile")
	fmt.Printf("--- config:\n%v\n\n", config)
}
