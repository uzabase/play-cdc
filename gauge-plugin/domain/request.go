package domain

import (
	"fmt"
	"sort"
	"strings"
)

type Request struct {
	Method      string         `json:"method"`
	Url         string         `json:"url"`
	UrlPath     string         `json:"urlPath"`
	QueryParams QueryParams    `json:"queryParameters"`
	Headers     RequestHeaders `json:"headers"`
}

func (r *Request) toRequestStep() Step {
	var request string
	if len(r.Headers) > 0 {
		request = fmt.Sprintf(`URL"%s"にヘッダー"%s"で%sリクエストを送る`, r.toUrl(), r.Headers, r.Method)
	} else {
		request = fmt.Sprintf(`URL"%s"に%sリクエストを送る`, r.toUrl(), r.Method)
	}
	return Step(request)
}

func (r *Request) toUrl() Step {
	var url string
	if len(r.Url) > 0 {
		url = r.Url
	} else {
		url = r.UrlPath + r.QueryParams.String()
	}
	return Step(url)
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
