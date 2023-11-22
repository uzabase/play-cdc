package domain

import (
	"fmt"
	"sort"
)

type Response struct {
	Status  int
	Headers ResponseHeaders
	Body    ResponseBody
}

type ResponseBody interface {
	toAssertions() []Step
}

type ResponseHeaders map[string]string

func (r *Response) toStatusCodeStep() Step {
	return Step(fmt.Sprintf(`レスポンスステータスコードが"%d"である`, r.Status))
}

func (r *Response) toHeaderAssertions() []Step {
	var assertions []Step
	for k, v := range r.Headers {
		assertions = append(assertions, Step(fmt.Sprintf(`レスポンスヘッダーに"%s"が存在し、その値が"%s"である`, k, v)))
	}

	sort.Slice(assertions, func(i, j int) bool { return assertions[i] < assertions[j] })

	return assertions
}

func (r *Response) toBodyAssertions() []Step {
	if r.Body == nil {
		return []Step{}
	}

	return r.Body.toAssertions()
}
