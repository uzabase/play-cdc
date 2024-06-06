package domain

import (
	"fmt"
)

type Contracts []*Contract

func (c Contracts) ToSpec(consumerName string, providerName string) *Spec {
	heading := toSpecHeading(consumerName, providerName)

	scenarios := make([]*Scenario, len(c))
	for i, contract := range c {
		scenarios[i] = contract.ToScenario(consumerName)
	}

	return NewSpec(heading, scenarios)
}

func toSpecHeading(consumerName string, providerName string) SpecHeading {
	return SpecHeading(fmt.Sprintf("%sが依存している%sの仕様", consumerName, providerName))
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
	base := fmt.Sprintf("%s %s", c.Request.Method, c.Request.toUrl())
	return ScenarioHeading(fmt.Sprintf("%s%s%s", base, c.Request.toHeaderHeading(), c.Request.toBodyHeading()))
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
