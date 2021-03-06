package usecase

import (
	"play-cdc/repository"
)

func GenerateSpec() {
	for _, e := range repository.GetEnvs() {
		requests, err := repository.FindExecutedRequests(e.APIEndpoint)
		if err != nil {
			panic(err)
		}

		contracts := requests.ToContracts()

		spec := contracts.ToSpec(e.APIName)
		spec.SortScenarios()

		repository.SaveSpec(spec, e.SpecPath)
	}
}
