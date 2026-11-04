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
	if len(stub.Url) > 0 {
		return Url{value: stub.Url}
	}
	if len(stub.UrlPath) > 0 {
		if query, ok := stub.QueryParams.toQueryString(); ok {
			return Url{value: stub.UrlPath + query}
		}
	}
	return Url{value: actual.Url}
}

func (u Url) String() string {
	return u.value
}

type QueryParams map[string]QueryParamMatcher

type QueryParamMatcher struct {
	EqualTo         *string `json:"equalTo"`
	Absent          bool    `json:"absent"`
	CaseInsensitive bool    `json:"caseInsensitive"`
}

// Conditions without a literal value must use the recorded request URL instead.
func (h QueryParams) toQueryString() (string, bool) {
	var queryParams []string
	for k, v := range h {
		if v.Absent {
			continue
		}
		if v.EqualTo == nil {
			return "", false
		}
		queryParams = append(queryParams, fmt.Sprintf(`%s=%s`, k, *v.EqualTo))
	}

	if len(queryParams) == 0 {
		return "", true
	}

	sort.Strings(queryParams)
	return "?" + strings.Join(queryParams, "&"), true
}
