package usecase

import (
	"specify/repository"
)

func GenerateSpec() {
	requests, err := repository.FindExecutedRequests(repository.WiremockEndpoint())
	if err != nil {
		panic(err)
	}

	contracts := requests.ToContracts()

	spec := contracts.ToSpec(repository.APIName())
	spec.SortScenarios()

	repository.SaveSpec(spec, repository.SpecPath())
}
