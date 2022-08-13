package domain

import (
	"fmt"
	"path/filepath"
	"strings"
)

func OutputPath(basePath string, providerName string) string {
	return filepath.Join(basePath, toPath(providerName))
}

func SpecFilePath(outputPath string, consumerName string) string {
	return filepath.Join(outputPath, fmt.Sprintf("%s.spec", toPath(consumerName)))
}

func RequestBodiesPath(outputPath string, consumerName string) string {
	return filepath.Join(outputPath, RequestBodiesRelativePath(consumerName))
}

func RequestBodiesRelativePath(consumerName string) string {
	return filepath.Join("fixtures", "contracts", toPath(consumerName))
}

func toPath(name string) string {
	lower := strings.ToLower(name)
	return strings.ReplaceAll(lower, " ", "_")
}
