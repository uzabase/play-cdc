package domain_test

import (
	"specify/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sut domain.Contract = domain.Contract{
	domain.Request{
		"/test",
		"GET",
	},
	domain.Response{
		200,
		nil,
	},
}

func TestToScenario_シナリオ名(t *testing.T) {
	actual := sut.ToScenario()

	assert.Equal(t, `GET /test`, actual.Heading)
}

func TestToScenario_リクエスト(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にGETリクエストを送る`))
}

func TestToScenario_レスポンスステータスコード(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスステータスコードが"200"である`))
}
