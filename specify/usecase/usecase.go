package usecase

import (
	"specify/repository"
)

const (
	wiremockEndpoint = "http://localhost:8080"
	specName         = "Example"
)

func GenerateSpec() {
	contracts, err := repository.FindExecutedContracts(wiremockEndpoint)
	if err != nil {
		panic(err)
	}

	spec := contracts.ToSpec(specName)

	repository.SaveSpec(spec)
}
