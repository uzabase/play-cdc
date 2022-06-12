package domain

type ExecutedRequests struct {
	Requests []ExecutedRequest `json:"requests"`
}

type ExecutedRequest struct {
	WasMatched  bool     `json:"wasMatched"`
	StubMapping Contract `json:"stubMapping"`
}
