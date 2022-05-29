package domain

import (
	"fmt"
)

type Contract struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

func (c *Contract) ToScenario() *Scenario {
	return &Scenario{
		Heading: c.toHeading(),
		Steps: c.toSteps(),
	}
}

func (c *Contract) toHeading() Heading {
	return Heading(fmt.Sprintf("%s %s", c.Request.Method, c.Request.toUrl()))
}

func (c *Contract) toSteps() []Step {
	steps := []Step{
		c.Request.toRequestStep(),
		c.Response.toStatusCodeStep(),
	}

	return append(steps, c.Response.toAssertions()...)
}
