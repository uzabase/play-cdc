package domain_test

import (
	"specify/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SutParams struct {
	method  string
	url     string
	urlPath string
	queryParams domain.QueryParams
	headers domain.Headers
}

var sut = createSut(SutParams{
	method:  "GET",
	urlPath: "/test",
})

func createSut(params SutParams) *domain.Contract {
	return &domain.Contract{
		Request: domain.Request{
			Url:     params.url,
			UrlPath: params.urlPath,
			Method:  params.method,
			QueryParams: params.queryParams,
			Headers: params.headers,
		},
		Response: domain.Response{
			Status: 200,
			JsonBody: map[string]any{
				"stringKey":  "stringValue",
				"integerKey": float64(123),
				"floatKey":   123.456,
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

func TestToScenario_シナリオ名にurlPathを使う(t *testing.T) {
	sut := createSut(SutParams{
		method:  "GET",
		urlPath: "/test",
	})

	actual := sut.ToScenario()

	assert.Equal(t, `GET /test`, actual.Heading)
}

func TestToScenario_シナリオ名にurlを使う(t *testing.T) {
	sut := createSut(SutParams{
		method:  "GET",
		url: "/test",
	})

	actual := sut.ToScenario()

	assert.Equal(t, `GET /test`, actual.Heading)
}

func TestToScenario_リクエストパスにurlPathを使う(t *testing.T) {
	sut := createSut(SutParams{
		method:  "GET",
		urlPath: "/test",
	})

	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にGETリクエストを送る`))
}

func TestToScenario_リクエストパスにurlを使う(t *testing.T) {
	sut := createSut(SutParams{
		method:  "GET",
		url: "/test",
	})

	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にGETリクエストを送る`))
}

func TestToScenario_GETリクエスト(t *testing.T) {
	sut := createSut(SutParams{
		method:  "GET",
		urlPath: "/test",
	})

	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にGETリクエストを送る`))
}

func TestToScenario_POSTリクエスト(t *testing.T) {
	sut := createSut(SutParams{
		method:  "POST",
		urlPath: "/test",
	})

	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にPOSTリクエストを送る`))
}

func TestToScenario_PUTリクエスト(t *testing.T) {
	sut := createSut(SutParams{
		method:  "PUT",
		urlPath: "/test",
	})

	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にPUTリクエストを送る`))
}

func TestToScenario_DELETEリクエスト(t *testing.T) {
	sut := createSut(SutParams{
		method:  "DELETE",
		urlPath: "/test",
	})

	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にDELETEリクエストを送る`))
}

func TestToScenario_クエリパラメータを含むリクエスト(t *testing.T) {
	sut := createSut(SutParams{
		method:  "GET",
		urlPath: "/test",
		queryParams: domain.QueryParams{
			"q1": {
				"equalTo": "v1",
			},
			"q2": {
				"equalTo": "v2",
			},
		},
	})

	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test?q1=v1&q2=v2"にGETリクエストを送る`))
}

func TestToScenario_ヘッダを含むリクエスト(t *testing.T) {
	sut := createSut(SutParams{
		method:  "PUT",
		urlPath: "/test",
		headers: domain.Headers{
			"content-type": {
				"equalTo": "application/json",
			},
			"options": {
				"equalTo": "123, 456",
			},
		},
	})

	actual := sut.ToScenario()

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にヘッダー"content-type: application/json \r\n options: 123, 456"でPUTリクエストを送る`))
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
