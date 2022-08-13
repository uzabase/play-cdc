package repository

import (
	"fmt"
	"os"
)

type Env struct {
	APIEndpoint string
	APIName     string
	OutputPath  string
}

func GetEnvs() []Env {
	var result []Env

	for i := 1; ; i++ {
		apiEndpoint := os.Getenv(fmt.Sprintf("cdc_api_endpoint_%d", i))
		apiName := os.Getenv(fmt.Sprintf("cdc_api_name_%d", i))
		outputPath := os.Getenv(fmt.Sprintf("cdc_output_path_%d", i))

		if apiEndpoint == "" || apiName == "" || outputPath == "" {
			break
		}

		result = append(result, Env{
			APIEndpoint: apiEndpoint,
			APIName:     apiName,
			OutputPath:  outputPath,
		})
	}

	return result
}
