package main

import (
	"fmt"
	"specify/repository"
	"specify/domain"
)

func main() {
	contracts := repository.FindExecutedContracts()

	scenarios := make([]*domain.Scenario, len(contracts))
	for i, contract := range contracts {
		scenarios[i] = contract.ToScenario()
	}

	spec := domain.NewSpec("Example", scenarios)
	fmt.Print(spec)
}
