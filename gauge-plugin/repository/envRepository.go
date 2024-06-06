package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"play-cdc/domain"
	"strconv"
)

func GetProviderEnvs() []domain.ProviderEnv {
	var result []domain.ProviderEnv

	for i := 1; ; i++ {
		providerEndpointKey := fmt.Sprintf("cdc_provider_endpoint_%d", i)
		providerNameKey := fmt.Sprintf("cdc_provider_name_%d", i)

		providerEndpoint := os.Getenv(providerEndpointKey)
		providerName := os.Getenv(providerNameKey)

		if providerEndpoint == "" && providerName == "" {
			break
		}

		if providerEndpoint == "" || providerName == "" {
			fmt.Println("Warning: properties are inconsistent.")
			fmt.Printf("%s=%s\n", providerEndpointKey, providerEndpoint)
			fmt.Printf("%s=%s\n", providerNameKey, providerName)
			break
		}

		result = append(result, domain.ProviderEnv{
			ProviderEndpoint: providerEndpoint,
			ProviderName:     providerName,
		})
	}

	return result
}

func GetOutputBasePath() string {
	outputBasePath := os.Getenv("cdc_output_base_path")

	if outputBasePath == "" {
		panic("error: output base path is empty!")
	}

	if !filepath.IsAbs(outputBasePath) {
		panic(fmt.Errorf("error: output base path is not an absolute path! current value: %s", outputBasePath))
	}

	return outputBasePath
}

func GetConsumerName() string {
	consumerName := os.Getenv("cdc_consumer_name")

	if consumerName == "" {
		panic("error: consumer name is empty!")
	}

	return consumerName
}

func IsDebug() bool {
	debug, err := strconv.ParseBool(os.Getenv("cdc_debug"))

	if err != nil {
		return false
	}

	return debug
}
