package main

import (
	"fmt"
	"specify/domain"
	"specify/repository"
)

const (
	wiremockEndpoint = "http://localhost:8080"
	specName         = "Example"
)

func main() {
	contracts := repository.FindExecutedContracts(wiremockEndpoint)

	scenarios := make([]*domain.Scenario, len(contracts))
	for i, contract := range contracts {
		scenarios[i] = contract.ToScenario()
	}

	spec := domain.NewSpec(specName, scenarios)
	fmt.Print(spec)
}
