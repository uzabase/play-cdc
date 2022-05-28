package domain

import (
	"fmt"
	"math"
	"strings"
)

type Contract struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Request struct {
	UrlPath string `json:"urlPath"`
	Method  string `json:"method"`
}

type Response struct {
	Status   int        `json:"status"`
	JsonBody map[string]any `json:"jsonBody"`
}

func (c Contract) ToScenario() Scenario {
	request := c.Request

	spec := Scenario{
		fmt.Sprintf("%s %s", request.Method, request.UrlPath),
		c.toSteps(),
	}

	return spec
}

func (c Contract) toSteps() []Step {
	request := fmt.Sprintf("URL\"%s\"に%sリクエストを送る", c.Request.UrlPath, c.Request.Method)
	statusCode := fmt.Sprintf("レスポンスステータスコードが\"%d\"である", c.Response.Status)

	steps := []Step{
		Step(request),
		Step(statusCode),
	}

	return append(steps, objectToAssertions(c.Response.JsonBody, []string{})...)
}

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
		step := fmt.Sprintf(`レスポンスのJSONの"$.%s"が文字列の"%s"である`, keyChain.toPath(), v)
		steps = []Step{Step(step)}
	case float64:
		steps = []Step{Step(toNumberStep(v, keyChain))}
	default:
		fmt.Printf("Warning: toAssertions - I don't know about type %T!\n", v)
	}

	return steps
}

func toNumberStep(n float64, keyChain KeyChain) string {
	if n == math.Trunc(n) {
		return fmt.Sprintf(`レスポンスのJSONの"$.%s"が整数の"%d"である`, keyChain.toPath(), int64(n))
	} else {
		return fmt.Sprintf(`レスポンスのJSONの"$.%s"が小数の"%g"である`, keyChain.toPath(), n)
	}
}
