package repository

import (
	"encoding/json"
	"io/ioutil"
	"specify/domain"
)

func FindContract(path string) *domain.Contract {
	byteArray, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var contract domain.Contract
	err = json.Unmarshal(byteArray, &contract)
	if err != nil {
		panic(err)
	}

	return &contract
}
