package domain

type Contract struct {
	Request Request `json:"request"`
	Response Response `json:"response"`
}

type Request struct {
	UrlPath string `json:"urlPath"`
	Method string `json:"method"`
}

type Response struct {
	Status int `json:"status"`
	JsonBody interface{} `json:"jsonBody"`
}
