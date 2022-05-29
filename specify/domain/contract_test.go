package domain_test

import (
	"specify/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sut = createSut("GET")

func createSut(method string) *domain.Contract {
	return &domain.Contract{
		domain.Request{
			"/test",
			method,
		},
		domain.Response{
			200,
			map[string]any{
				"stringKey": "stringValue",
				"integerKey": float64(123),
				"floatKey": 123.456,
				"booleanKey": true,
				"objectKey": map[string]any{
					"stringKey": "objectStringValue",
				},
				"arrayKey": []any{
					map[string]any{
						"stringKey": "arrayObjectStringValue",
					},
				},
			},
		},
	}
}

func TestToScenario_シナリオ名(t *testing.T) {
	actual := sut.ToScenario()

	assert.Equal(t, `GET /test`, actual.Heading)
}

func TestToScenario_GETリクエスト(t *testing.T) {
	actual := createSut("GET").ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にGETリクエストを送る`))
}

func TestToScenario_POSTリクエスト(t *testing.T) {
	actual := createSut("POST").ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にPOSTリクエストを送る`))
}

func TestToScenario_PUTリクエスト(t *testing.T) {
	actual := createSut("PUT").ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にPUTリクエストを送る`))
}

func TestToScenario_DELETEリクエスト(t *testing.T) {
	actual := createSut("DELETE").ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にDELETEリクエストを送る`))
}

func TestToScenario_レスポンスステータスコード(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスステータスコードが"200"である`))
}

func TestToScenario_文字列のアサーション(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.stringKey"が文字列の"stringValue"である`))
}

func TestToScenario_整数のアサーション(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.integerKey"が整数の"123"である`))
}

func TestToScenario_小数のアサーション(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.floatKey"が小数の"123.456"である`))
}

func TestToScenario_真偽値のアサーション(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.booleanKey"が真偽値の"true"である`))
}

func TestToScenario_オブジェクトに含まれる値のアサーション(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.objectKey.stringKey"が文字列の"objectStringValue"である`))
}

func TestToScenario_配列に含まれるオブジェクトに含まれる値のアサーション(t *testing.T) {
	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.arrayKey[0].stringKey"が文字列の"arrayObjectStringValue"である`))
}
