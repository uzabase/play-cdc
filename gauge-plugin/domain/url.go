package domain

import (
	"fmt"
	"sort"
	"strings"
)

type Url struct {
	value string
}

func NewUrl(stub StubRequest, actual ActualRequest) Url {
	var value string
	if len(stub.Url) > 0 {
		value = stub.Url
	} else if len(stub.UrlPath) > 0 {
		value = stub.UrlPath + stub.QueryParams.String()
	} else {
		value = actual.Url
	}
	return Url{value: value}
}

func (u Url) String() string {
	return u.value
}

type QueryParams map[string]QueryParamMatcher

type QueryParamMatcher map[string]string

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
