package domain

import (
	"fmt"
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
	Status   int  `json:"status"`
	JsonBody JSON `json:"jsonBody"`
}

type JSON map[string]interface{}
type JSONArray []interface{}

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

	return append(steps, c.Response.JsonBody.toAssertions([]string{})...)
}

type KeyChain []string

func (k KeyChain) toPath() string {
	return strings.Join(k, ".")
}

func (j JSON) toAssertions(keyChain KeyChain) []Step {
	var assertions []Step
	for k, v := range j {
		keyChain := append(keyChain, k)
		assertions = append(assertions, toAssertions(v, keyChain)...)
	}
	return assertions
}

func (j JSONArray) toAssertions(keyChain KeyChain) []Step {
	var assertions []Step
	for i, v := range j {
		keyChain[len(keyChain) - 1] += fmt.Sprintf("[%d]", i)
		assertions = append(assertions, toAssertions(v, keyChain)...)
	}
	return assertions
}

func toAssertions(v interface{}, keyChain KeyChain) []Step {
	var steps []Step

	switch v := v.(type) {
	case map[string]interface{}:
		steps = JSON(v).toAssertions(keyChain)
	case []interface{}:
		steps = JSONArray(v).toAssertions(keyChain)
	case string:
		step := fmt.Sprintf(`レスポンスのJSONの"$.%s"が文字列の"%s"である`, keyChain.toPath(), v)
		steps = []Step{Step(step)}
	default:
		fmt.Printf("Warning: toAssertions - I don't know about type %T!\n", v)
	}

	return steps
}
