package main

import (
	"fmt"
	"specify/repository"
)

func main() {
	contracts := repository.FindExecutedContracts()

	for _, contract := range contracts {
		scenario := contract.ToScenario()
		fmt.Println(scenario)
	}
}
