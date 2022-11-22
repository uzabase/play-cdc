package domain

import (
	"encoding/json"
	"reflect"
)

type ExecutedRequests struct {
	Requests []ExecutedRequest `json:"requests"`
}

type ExecutedRequest struct {
	Request     ActualRequest `json:"request"`
	WasMatched  bool          `json:"wasMatched"`
	StubMapping StubMapping   `json:"stubMapping"`
}

type ActualRequest struct {
	Body string `json:"body"`
}

type StubMapping struct {
	Request  StubRequest  `json:"request"`
	Response StubResponse `json:"response"`
}
type StubRequest struct {
	Method      string         `json:"method"`
	Url         string         `json:"url"`
	UrlPath     string         `json:"urlPath"`
	QueryParams QueryParams    `json:"queryParameters"`
	Headers     RequestHeaders `json:"headers"`
}

type StubResponse struct {
	Status   int             `json:"status"`
	Headers  ResponseHeaders `json:"headers"`
	Body     string          `json:"body"`
	JsonBody map[string]any  `json:"jsonBody"`
}

func ToContracts(requests []ExecutedRequest) Contracts {
	var result Contracts
	for _, r := range requests {
		s := r.StubMapping

		if !r.WasMatched {
			continue
		}

		c := &Contract{
			Request: Request{
				Method:      s.Request.Method,
				Url:         s.Request.Url,
				UrlPath:     s.Request.UrlPath,
				QueryParams: s.Request.QueryParams,
				Headers:     s.Request.Headers,
				Body:        r.Request.Body,
			},
			Response: Response{
				Status: s.Response.Status,
				Headers: s.Response.Headers,
				JsonBody: s.Response.toJsonBody(),
			},
		}

		if result.contains(c) {
			continue
		}

		result = append(result, c)
	}

	return result
}

func (r *StubResponse) toJsonBody() (map[string]any) {
	if r.JsonBody != nil {
		return r.JsonBody
	}

	if len(r.Body) > 0 {
		var jsonBody map[string]any
		err := json.Unmarshal([]byte(r.Body), &jsonBody)
		if err != nil {
			panic(err)
		}
		return jsonBody
	}

	return nil
}

func (c Contracts) contains(contract *Contract) bool {
	for _, v := range c {
		if reflect.DeepEqual(v, contract) {
			return true
		}
	}
	return false
}
