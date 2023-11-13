package domain

import "fmt"

type TextResponseBody string

func CreateTextResponseBody(value string) ResponseBody {
	result := TextResponseBody(value)
	return &result
}

func (t *TextResponseBody) toAssertions() []Step {
	return []Step{
		Step(fmt.Sprintf(`レスポンスボディが文字列"%s"である`, string(*t))),
	}
}
