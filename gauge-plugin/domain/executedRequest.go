package domain

type ExecutedRequests struct {
	Requests []ExecutedRequest `json:"requests"`
}

type ExecutedRequest struct {
	Request     Request     `json:"request"`
	WasMatched  bool        `json:"wasMatched"`
	StubMapping StubMapping `json:"stubMapping"`
}

type Request struct {
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
			Request:  s.Request,
			Response: s.Response,
		})
		added[s.Id] = struct{}{}
	}

	return result
}
