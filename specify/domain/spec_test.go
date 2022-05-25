package domain_test

import (
	"specify/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	sut := domain.Scenario{
		"heading",
		[]domain.Step{
			"step1",
			"step2",
		},
	}

	actual := sut.String()

	expected := `## heading
* step1
* step2
`
	assert.Equal(t, expected, actual)
}
