package domain_test

import (
	"play-cdc/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputPath(t *testing.T) {
	assert.Equal(t, "/path/provider_api", domain.OutputPath("/path", "Provider API"))
}

func TestSpecFilePath(t *testing.T) {
	assert.Equal(t, "/path/provider_api/consumer_api.spec", domain.SpecFilePath("/path/provider_api", "Consumer API"))
}

func TestRequestBodiesPath(t *testing.T) {
	assert.Equal(t, "/path/provider_api/fixtures/contracts/consumer_api", domain.RequestBodiesPath("/path/provider_api", "Consumer API"))
}

func TestRequestBodiesRelativePath(t *testing.T) {
	assert.Equal(t, "fixtures/contracts/consumer_api", domain.RequestBodiesRelativePath("Consumer API"))
}
