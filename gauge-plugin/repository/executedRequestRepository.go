package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"play-cdc/domain"
)

var recordedRequests = map[string][]domain.ExecutedRequest{}

func FetchExecutedRequests(endpoint string) (*domain.ExecutedRequests, error) {
	req, err := http.NewRequest("GET", endpoint+"/__admin/requests", nil)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to create a request for endpoint(%s): %w", endpoint, err)
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to request for endpoint(%s): %w", endpoint, err)
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to read response body from endpoint(%s): %w", endpoint, err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error: received non-ok response from endpoint(%s). The response body was: %s", endpoint, byteArray)
	}

	result, err := Unmarshal(byteArray)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to unmarshal json read from endpoint(%s): %w", endpoint, err)
	}

	return result, nil
}

func RecordExecutedRequests(key string, requests []domain.ExecutedRequest) {
	recordedRequests[key] = append(recordedRequests[key], requests...)
}

func LoadRecordedRequests(key string) []domain.ExecutedRequest {
	return recordedRequests[key]
}

func ClearExecutedRequests() {
	recordedRequests = map[string][]domain.ExecutedRequest{}
}

func Unmarshal(byteArray []byte) (*domain.ExecutedRequests, error) {
	var result domain.ExecutedRequests
	err := json.Unmarshal(byteArray, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}
