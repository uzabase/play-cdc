package main

import (
	"fmt"
	"specify/domain"
	"specify/handler"
	"specify/repository"
)

const (
	wiremockEndpoint = "http://localhost:8080"
	specName         = "Example"
)

func main() {
	handler.Start()

	contracts, err := repository.FindExecutedContracts(wiremockEndpoint)
	if err != nil {
		panic(err)
	}

	scenarios := make([]*domain.Scenario, len(contracts))
	for i, contract := range contracts {
		scenarios[i] = contract.ToScenario()
	}

	spec := domain.NewSpec(specName, scenarios)
	fmt.Print(spec)
}
