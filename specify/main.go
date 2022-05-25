package main

import (
	"fmt"
	"specify/repository"
)

func main() {
	contract := repository.FindContract("example.json")
	scenario := contract.ToScenario()
	fmt.Println(scenario)
}
