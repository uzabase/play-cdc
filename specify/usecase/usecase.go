package usecase

import (
	"specify/repository"
)

const (
	wiremockEndpoint = "http://localhost:16000"
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
