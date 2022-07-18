package domain

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Response struct {
	Status   int             `json:"status"`
	Headers  ResponseHeaders `json:"headers"`
	JsonBody map[string]any  `json:"jsonBody"`
}

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
	return objectToAssertions(r.JsonBody, []string{})
}

type ResponseHeaders map[string]string

type KeyChain []string

func (k KeyChain) toPath() string {
	return strings.Join(k, ".")
}

func objectToAssertions(object map[string]any, keyChain KeyChain) []Step {
	var assertions []Step
	for k, v := range object {
		keyChain := append(keyChain, k)
		assertions = append(assertions, toAssertions(v, keyChain)...)
	}

	sort.Slice(assertions, func(i, j int) bool { return assertions[i] < assertions[j] })
	return assertions
}

func arrayToAssertions(array []any, keyChain KeyChain) []Step {
	var assertions []Step
	for i, v := range array {
		keyChain[len(keyChain)-1] += fmt.Sprintf("[%d]", i)
		assertions = append(assertions, toAssertions(v, keyChain)...)
	}
	return assertions
}

func toAssertions(v any, keyChain KeyChain) []Step {
	var steps []Step

	switch v := v.(type) {
	case map[string]any:
		steps = objectToAssertions(v, keyChain)
	case []any:
		steps = arrayToAssertions(v, keyChain)
	case string:
		steps = []Step{Step(toStringStep(v, keyChain))}
	case float64:
		steps = []Step{Step(toNumberStep(v, keyChain))}
	case bool:
		steps = []Step{Step(toBoolStep(v, keyChain))}
	default:
		fmt.Printf("Warning: toAssertions - I don't know about type %T!\n", v)
	}

	return steps
}

func toStringStep(s string, keyChain KeyChain) string {
	return fmt.Sprintf(`レスポンスのJSONの"$.%s"が文字列の"%s"である`, keyChain.toPath(), s)
}

func toNumberStep(n float64, keyChain KeyChain) string {
	if n == math.Trunc(n) {
		return fmt.Sprintf(`レスポンスのJSONの"$.%s"が整数の"%d"である`, keyChain.toPath(), int64(n))
	} else {
		return fmt.Sprintf(`レスポンスのJSONの"$.%s"が小数の"%g"である`, keyChain.toPath(), n)
	}
}

func toBoolStep(b bool, keyChain KeyChain) string {
	return fmt.Sprintf(`レスポンスのJSONの"$.%s"が真偽値の"%t"である`, keyChain.toPath(), b)
}
