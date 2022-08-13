package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"play-cdc/domain"
)

func FindExecutedRequests(endpoint string) (*domain.ExecutedRequests, error) {
	req, err := http.NewRequest("GET", endpoint+"/__admin/requests", nil)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to create a request for endpoint(%s): %w", endpoint, err)
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to request for endpoint(%s): %w", endpoint, err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to read response body from endpoint(%s): %w", endpoint, err)
	}

	var result domain.ExecutedRequests
	err = json.Unmarshal(byteArray, &result)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to unmarshal json read from endpoint(%s): %w", endpoint, err)
	}

	return &result, nil
}
