package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"specify/domain"
)

func FindContract(path string) (*domain.Contract, error) {
	byteArray, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file from path(%s): %w", path, err)
	}

	var contract domain.Contract
	err = json.Unmarshal(byteArray, &contract)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal json read from path(%s): %w", path, err)
	}

	return &contract, nil
}
