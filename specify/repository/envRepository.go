package repository

import (
	"fmt"
	"os"
)

type Env struct {
	APIEndpoint string
	APIName     string
	SpecPath    string
}

func GetEnvs() []Env {
	var result []Env

	for i := 1; ; i++ {
		apiEndpoint := os.Getenv(fmt.Sprintf("cdc_api_endpoint_%d", i))
		apiName := os.Getenv(fmt.Sprintf("cdc_api_name_%d", i))
		specPath := os.Getenv(fmt.Sprintf("cdc_spec_path_%d", i))

		if apiEndpoint == "" || apiName == "" || specPath == "" {
			break
		}

		result = append(result, Env{
			APIEndpoint: apiEndpoint,
			APIName:     apiName,
			SpecPath:    specPath,
		})
	}

	return result
}
