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
		apiEndpointKey := fmt.Sprintf("cdc_api_endpoint_%d", i)
		apiNameKey := fmt.Sprintf("cdc_api_name_%d", i)
		outputPathKey := fmt.Sprintf("cdc_output_path_%d", i)

		apiEndpoint := os.Getenv(apiEndpointKey)
		apiName := os.Getenv(apiNameKey)
		outputPath := os.Getenv(outputPathKey)

		if apiEndpoint == "" && apiName == "" && outputPath == "" {
			break
		}

		if apiEndpoint == "" || apiName == "" || outputPath == "" {
			fmt.Println("Warning: properties are inconsistent.")
			fmt.Printf("%s=%s\n", apiEndpointKey, apiEndpoint)
			fmt.Printf("%s=%s\n", apiNameKey, apiName)
			fmt.Printf("%s=%s\n", outputPathKey, outputPath)
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
