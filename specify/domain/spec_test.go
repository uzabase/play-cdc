package domain_test

import (
	"specify/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecString(t *testing.T) {
	scenario1 := &domain.Scenario{
		Heading: "scenario1",
		Steps: []domain.Step{
			"step1",
		},
	}

	scenario2 := &domain.Scenario{
		Heading: "scenario2",
		Steps: []domain.Step{
			"step1",
		},
	}

	sut := domain.NewSpec("spec", []*domain.Scenario{scenario1, scenario2})

	actual := sut.String()

	expected := `# spec

## scenario1
* step1

## scenario2
* step1
`
	assert.Equal(t, expected, actual)
}

func TestScenarioString(t *testing.T) {
	sut := domain.Scenario{
		Heading: "heading",
		Steps: []domain.Step{
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
