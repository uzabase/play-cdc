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
		domain.JSON{
			"stringKey": "stringValue",
			"objectKey": domain.JSON{
				"stringKey": "objectStringValue",
			},
		},
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

func TestToScenario_文字列のアサーション(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.stringKey"が文字列の"stringValue"である`))
}

func TestToScenario_オブジェクトに含まれる文字列のアサーション(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.objectKey.stringKey"が文字列の"objectStringValue"である`))
}
