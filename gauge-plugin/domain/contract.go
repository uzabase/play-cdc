package domain

import (
	"fmt"
)

type Contracts []*Contract

func (c Contracts) ToSpec(consumerName string, providerName string) *Spec {
	scenarios := make([]*Scenario, len(c))

	for i, contract := range c {
		scenarios[i] = contract.ToScenario(consumerName)
	}

	return NewSpec(toSpecHeading(consumerName, providerName), scenarios)
}

func toSpecHeading(consumerName string, providerName string) string {
	return fmt.Sprintf("%sが依存している%sの仕様", consumerName, providerName)
}

type Contract struct {
	Request  Request
	Response Response
}

func (c *Contract) ToScenario(consumerName string) *Scenario {
	return &Scenario{
		Heading: c.toHeading(),
		Steps:   c.toSteps(consumerName),
	}
}

func (c *Contract) toHeading() ScenarioHeading {
	if c.Request.IsBodyAvailable() {
		return ScenarioHeading(fmt.Sprintf("%s %s (body: %x)", c.Request.Method, c.Request.toUrl(), c.Request.toBodyHash()))
	} else {
		return ScenarioHeading(fmt.Sprintf("%s %s", c.Request.Method, c.Request.toUrl()))
	}
}

func (c *Contract) toSteps(consumerName string) []Step {
	steps := []Step{
		c.Request.toRequestStep(consumerName),
		c.Response.toStatusCodeStep(),
	}

	steps = append(steps, c.Response.toHeaderAssertions()...)
	steps = append(steps, c.Response.toBodyAssertions()...)
	return steps
}
