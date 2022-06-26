package usecase

import (
	"specify/repository"
)

const (
	wiremockEndpoint = "http://localhost:16000"
	specName         = "Example"
)

func GenerateSpec() {
	requests, err := repository.FindExecutedRequests(wiremockEndpoint)
	if err != nil {
		panic(err)
	}

	contracts := requests.ToContracts()
	spec := contracts.ToSpec(specName)

	repository.SaveSpec(spec)
}
