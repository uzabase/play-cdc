package repository

import (
	"fmt"
	"os"
	"strconv"
)

type ProviderEnv struct {
	ProviderEndpoint string
	ProviderName     string
}

func GetProviderEnvs() []ProviderEnv {
	var result []ProviderEnv

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

		result = append(result, ProviderEnv{
			ProviderEndpoint: providerEndpoint,
			ProviderName:     providerName,
		})
	}

	return result
}

func GetOutputBasePath() string {
	return os.Getenv("cdc_output_base_path")
}

func GetConsumerName() string {
	return os.Getenv("cdc_consumer_name")
}

func IsDebug() bool {
	debug, err := strconv.ParseBool(os.Getenv("cdc_debug"))
	if err != nil {
		return false
	}
	return debug
}
