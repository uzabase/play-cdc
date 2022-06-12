package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"specify/domain"
)

func FindExecutedContracts(endpoint string) []domain.Contract {
	req, err := http.NewRequest("GET", endpoint+"/__admin/requests", nil)
	if err != nil {
		panic(err)
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var executedRequests domain.ExecutedRequests
	err = json.Unmarshal(byteArray, &executedRequests)
	if err != nil {
		panic(err)
	}

	var result []domain.Contract
	for _, r := range executedRequests.Requests {
		if r.WasMatched {
			result = append(result, r.StubMapping)
		}
	}

	return result
}
