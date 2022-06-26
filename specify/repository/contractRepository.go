package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"specify/domain"
)

func FindExecutedContracts(endpoint string) (domain.Contracts, error) {
	req, err := http.NewRequest("GET", endpoint+"/__admin/requests", nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a request for endpoint(%s): %w", endpoint, err)
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to request for endpoint(%s): %w", endpoint, err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body from endpoint(%s): %w", endpoint, err)
	}

	var executedRequests domain.ExecutedRequests
	err = json.Unmarshal(byteArray, &executedRequests)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal json read from endpoint(%s): %w", endpoint, err)
	}

	var result []*domain.Contract
	for _, r := range executedRequests.Requests {
		if r.WasMatched {
			result = append(result, &r.StubMapping)
		}
	}

	return result, nil
}
