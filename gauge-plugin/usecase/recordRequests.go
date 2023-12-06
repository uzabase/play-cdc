package usecase

import (
	"play-cdc/repository"
)

func RecordRequests() {
	for _, e := range repository.GetProviderEnvs() {
		requests, err := repository.FetchExecutedRequests(e.ProviderEndpoint)
		if err != nil {
			panic(err)
		}

		repository.RecordExecutedRequests(e.ProviderEndpoint, requests.Requests)
	}
}
