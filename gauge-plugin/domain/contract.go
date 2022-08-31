package domain

import (
	"crypto/md5"
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
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

func (r *Request) toRequestStep(consumerName string) Step {
	var request string
	if len(r.Headers) > 0 {
		if r.IsBodyAvailable() {
			request = fmt.Sprintf(`URL"%s"にボディ"<file:%s>"、ヘッダー"%s"で、%sリクエストを送る`, r.toUrl(), r.bodyFilePath(consumerName), r.Headers, r.Method)
		} else {
			request = fmt.Sprintf(`URL"%s"にヘッダー"%s"で、%sリクエストを送る`, r.toUrl(), r.Headers, r.Method)
		}
	} else {
		if r.IsBodyAvailable() {
			request = fmt.Sprintf(`URL"%s"にボディ"<file:%s>"で、%sリクエストを送る`, r.toUrl(), r.bodyFilePath(consumerName), r.Method)
		} else {
			request = fmt.Sprintf(`URL"%s"に%sリクエストを送る`, r.toUrl(), r.Method)
		}
	}
	return Step(request)
}

func (r *Request) bodyFilePath(consumerName string) string {
	return filepath.Join(RequestBodiesRelativePath(consumerName), r.ToBodyFileName())
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

func (r *Request) toBodyHash() [16]byte {
	return md5.Sum([]byte(r.Body))
}

func (r *Request) ToBodyFileName() string {
	re := regexp.MustCompile("[/|?|=|&]")
	replaced := re.ReplaceAllString(r.toUrl()[1:], "_")
	return fmt.Sprintf("%s_%s_%x.json", strings.ToLower(r.Method), replaced, r.toBodyHash())
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
