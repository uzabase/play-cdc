package domain

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
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
	Request  Request
	Response StubResponse
}

type Request struct {
	Method      string
	Url         string
	UrlPath     string
	QueryParams QueryParams
	Headers     RequestHeaders
	Body        string
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

func (r *Request) toRequestStep() Step {
	var request string
	if len(r.Headers) > 0 {
		if r.IsBodyAvailable() {
			request = fmt.Sprintf(`URL"%s"にボディ"file:fixtures/%s"、ヘッダー"%s"で、%sリクエストを送る`, r.toUrl(), r.ToBodyFileName(), r.Headers, r.Method)
		} else {
			request = fmt.Sprintf(`URL"%s"にヘッダー"%s"で、%sリクエストを送る`, r.toUrl(), r.Headers, r.Method)
		}
	} else {
		if r.IsBodyAvailable() {
			request = fmt.Sprintf(`URL"%s"にボディ"file:fixtures/%s"で、%sリクエストを送る`, r.toUrl(), r.ToBodyFileName(), r.Method)
		} else {
			request = fmt.Sprintf(`URL"%s"に%sリクエストを送る`, r.toUrl(), r.Method)
		}
	}
	return Step(request)
}

func (r *Request) IsBodyAvailable() bool {
	return len(r.Body) > 0 && (r.Method == "POST" || r.Method == "PUT")
}

func (r *Request) toUrl() string {
	var url string
	if len(r.Url) > 0 {
		url = r.Url
	} else {
		url = r.UrlPath + r.QueryParams.String()
	}
	return url
}

func (r *Request) ToBodyFileName() string {
	re := regexp.MustCompile("[/|?|=|&]")
	replaced := re.ReplaceAllString(r.toUrl()[1:], "_")
	return replaced + ".json"
}

type QueryParams map[string]QueryParamMatcher

type QueryParamMatcher map[string]string

type RequestHeaders map[string]HeaderMather

type HeaderMather map[string]string

func (h QueryParams) String() string {
	if len(h) == 0 {
		return ""
	}

	var queryParams []string
	for k, v := range h {
		queryParams = append(queryParams, fmt.Sprintf(`%s=%s`, k, v))
	}

	sort.Slice(queryParams, func(i, j int) bool { return queryParams[i] < queryParams[j] })

	return "?" + strings.Join(queryParams, "&")
}

func (m QueryParamMatcher) String() string {
	return m["equalTo"]
}

func (h RequestHeaders) String() string {
	var headers []string
	for k, v := range h {
		headers = append(headers, fmt.Sprintf(`%s: %s`, k, v))
	}

	sort.Slice(headers, func(i, j int) bool { return headers[i] < headers[j] })

	return strings.Join(headers, ` \r\n `)
}

func (m HeaderMather) String() string {
	return m["equalTo"]
}
