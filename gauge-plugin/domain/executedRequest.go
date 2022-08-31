package domain

import "reflect"

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

func (er *ExecutedRequests) ToContracts() Contracts {
	var result Contracts
	for _, r := range er.Requests {
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
			Response: s.Response,
		}

		if exists(result, c) {
			continue
		}

		result = append(result, c)
	}

	return result
}

func exists(contracts Contracts, contract *Contract) bool {
	for _, v := range contracts {
		if reflect.DeepEqual(v, contract) {
			return true
		}
	}
	return false
}
