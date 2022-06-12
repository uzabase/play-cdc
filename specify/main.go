package main

import (
	"fmt"
	"os"
	"specify/domain"
	"specify/repository"
)

const (
	wiremockEndpoint = "http://localhost:8080"
	specName         = "Example"
)

func main() {
	contracts, err := repository.FindExecutedContracts(wiremockEndpoint)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	scenarios := make([]*domain.Scenario, len(contracts))
	for i, contract := range contracts {
		scenarios[i] = contract.ToScenario()
	}

	spec := domain.NewSpec(specName, scenarios)
	fmt.Print(spec)
}
