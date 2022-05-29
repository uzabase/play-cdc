package domain

import (
	"fmt"
	"strings"
)

type Scenario struct {
	Heading Heading
	Steps []Step
}

type Heading string

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
