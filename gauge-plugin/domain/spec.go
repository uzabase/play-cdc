package domain

import (
	"fmt"
	"sort"
	"strings"
)

type Spec struct {
	Heading   SpecHeading
	Scenarios []*Scenario
}

type SpecHeading string

func NewSpec(heading SpecHeading, scenarios []*Scenario) *Spec {
	return &Spec{
		Heading:   heading,
		Scenarios: scenarios,
	}
}

func (s *Spec) SortScenarios() {
	sort.Slice(s.Scenarios, func(i, j int) bool {
		return s.Scenarios[i].Heading < s.Scenarios[j].Heading
	})
}

func (s *Spec) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("# %s\n", s.Heading))
	b.WriteString("\n")

	scenarios := make([]string, len(s.Scenarios))
	for i, scenario := range s.Scenarios {
		scenarios[i] = scenario.String()
	}
	b.WriteString(strings.Join(scenarios, "\n"))

	return b.String()
}

type Scenario struct {
	Heading ScenarioHeading
	Steps   []Step
}

type ScenarioHeading string

type Step string

func (s *Scenario) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("## %s\n", s.Heading))

	for _, step := range s.Steps {
		b.WriteString(fmt.Sprintf("* %s\n", step.String()))
	}

	return b.String()
}

func (s Step) String() string {
	return string(s)
}
