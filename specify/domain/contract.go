package domain

import "fmt"

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

func (c Contract) ToScenario() Scenario {
	request := c.Request

	spec := Scenario{
		fmt.Sprintf("%s %s", request.Method, request.UrlPath),
		c.toSteps(),
	}

	return spec
}

func (c Contract) toSteps() []Step {
	request := fmt.Sprintf("URL\"%s\"に%sリクエストを送る", c.Request.UrlPath, c.Request.Method)
	statusCode := fmt.Sprintf("レスポンスステータスコードが\"%d\"である", c.Response.Status)

	return []Step{
		Step(request),
		Step(statusCode),
	}
}
