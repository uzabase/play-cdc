package usecase

import (
	"path/filepath"
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

		err = repository.SaveSpec(spec, filepath.Join(e.OutputPath, "contract.spec"))
		if err != nil {
			panic(err)
		}

		repository.SaveRequestBodies(contracts, e.OutputPath)
	}
}
