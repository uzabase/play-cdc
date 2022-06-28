package repository

import "os"

func WiremockEndpoint() string {
	return os.Getenv("wiremock_endpoint")
}

func APIName() string {
	return os.Getenv("api_name")
}

func SpecPath() string {
	return os.Getenv("spec_path")
}
