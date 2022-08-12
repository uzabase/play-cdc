package domain

type StubRequest struct {
	Method      string         `json:"method"`
	Url         string         `json:"url"`
	UrlPath     string         `json:"urlPath"`
	QueryParams QueryParams    `json:"queryParameters"`
	Headers     RequestHeaders `json:"headers"`
}
