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
	} else if len(stub.UrlPath) > 0 && stub.QueryParams.convertibleToRequest() {
		value = stub.UrlPath + stub.QueryParams.toQueryString()
	} else {
		value = actual.Url
	}
	return Url{value: value}
}

func (u Url) String() string {
	return u.value
}

type QueryParams map[string]QueryParamMatcher

type QueryParamMatcher struct {
	EqualTo *string `json:"equalTo"`
}

func (h QueryParams) convertibleToRequest() bool {
	for _, v := range h {
		if v.EqualTo == nil {
			return false
		}
	}
	return true
}

func (h QueryParams) toQueryString() string {
	var queryParams []string
	for k, v := range h {
		if v.EqualTo == nil {
			continue
		}
		queryParams = append(queryParams, fmt.Sprintf(`%s=%s`, k, *v.EqualTo))
	}
	if len(queryParams) == 0 {
		return ""
	}

	sort.Strings(queryParams)
	return "?" + strings.Join(queryParams, "&")
}
