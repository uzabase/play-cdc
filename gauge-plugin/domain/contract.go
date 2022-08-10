package domain

import (
	"fmt"
)

type Contracts []*Contract

func (c Contracts) ToSpec(specName string) *Spec {
	scenarios := make([]*Scenario, len(c))

	for i, contract := range c {
		scenarios[i] = contract.ToScenario()
	}

	return NewSpec(specName, scenarios)
}

type Contract struct {
	Request  StubRequest  `json:"request"`
	Response StubResponse `json:"response"`
}

func (c *Contract) ToScenario() *Scenario {
	return &Scenario{
		Heading: c.toHeading(),
		Steps:   c.toSteps(),
	}
}

func (c *Contract) toHeading() ScenarioHeading {
	return ScenarioHeading(fmt.Sprintf("%s %s", c.Request.Method, c.Request.toUrl()))
}

func (c *Contract) toSteps() []Step {
	steps := []Step{
		c.Request.toRequestStep(),
		c.Response.toStatusCodeStep(),
	}

	steps = append(steps, c.Response.toHeaderAssertions()...)
	steps = append(steps, c.Response.toBodyAssertions()...)
	return steps
}
