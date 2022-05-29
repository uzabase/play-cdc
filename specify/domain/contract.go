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
	Method  string `json:"method"`
	Url string `json:"url"`
	UrlPath string `json:"urlPath"`
	QueryParams QueryParams `json:"queryParameters"`
	Headers Headers `json:"headers"`
}

func (r *Request) toUrl() string {
	if (len(r.Url) > 0) {
		return r.Url
	}
	return r.UrlPath + r.QueryParams.String()
}

type QueryParams map[string]QueryParamMatcher

type QueryParamMatcher map[string]string

type Headers map[string]HeaderMather

type HeaderMather map[string]string

func (h QueryParams) String() string {
	if (len(h) == 0) {
		return ""
	}

	var result []string
	for k, v := range h {
		result = append(result, fmt.Sprintf(`%s=%s`, k, v))
	}
	return "?" + strings.Join(result, "&")
}

func (m QueryParamMatcher) String() string {
	return m["equalTo"]
}

func (h Headers) String() string {
	var result []string
	for k, v := range h {
		result = append(result, fmt.Sprintf(`%s: %s`, k, v))
	}
	return strings.Join(result, ` \r\n `)
}

func (m HeaderMather) String() string {
	return m["equalTo"]
}

type Response struct {
	Status   int        `json:"status"`
	JsonBody map[string]any `json:"jsonBody"`
}

func (c *Contract) ToScenario() *Scenario {
	request := c.Request

	return &Scenario{
		Heading: fmt.Sprintf("%s %s", request.Method, request.toUrl()),
		Steps: c.toSteps(),
	}
}

func (c *Contract) toSteps() []Step {
	steps := []Step{
		Step(c.Request.toRequestStep()),
		Step(c.Response.toStatusCodeStep()),
	}

	return append(steps, objectToAssertions(c.Response.JsonBody, []string{})...)
}

func (r *Request) toRequestStep() string {
	if (len(r.Headers) > 0) {
		return fmt.Sprintf(`URL"%s"にヘッダー"%s"で%sリクエストを送る`, r.toUrl(), r.Headers, r.Method)
	}
	return fmt.Sprintf(`URL"%s"に%sリクエストを送る`, r.toUrl(), r.Method)
}

func (r *Response) toStatusCodeStep() string {
	return fmt.Sprintf(`レスポンスステータスコードが"%d"である`, r.Status)
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
