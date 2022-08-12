package domain

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
	Id       string       `json:"id"`
	Request  StubRequest  `json:"request"`
	Response StubResponse `json:"response"`
}

func (er *ExecutedRequests) ToContracts() Contracts {
	var result Contracts
	var added = map[string]struct{}{}
	for _, r := range er.Requests {
		s := r.StubMapping

		if !r.WasMatched {
			continue
		}
		_, exists := added[s.Id]
		if exists {
			continue
		}

		result = append(result, &Contract{
			Request: Request{
				Method:      s.Request.Method,
				Url:         s.Request.Url,
				UrlPath:     s.Request.UrlPath,
				QueryParams: s.Request.QueryParams,
				Headers:     s.Request.Headers,
				Body:        r.Request.Body,
			},
			Response: s.Response,
		})
		added[s.Id] = struct{}{}
	}

	return result
}
