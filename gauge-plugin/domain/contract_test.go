package domain_test

import (
	"play-cdc/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SutParams struct {
	method      string
	url         string
	urlPath     string
	queryParams domain.QueryParams
	headers     domain.RequestHeaders
	body        string
}

var contract = createContract(SutParams{
	method:  "GET",
	urlPath: "/test",
})

func createContract(params SutParams) *domain.Contract {
	return &domain.Contract{
		Request: domain.Request{
			Url:         params.url,
			UrlPath:     params.urlPath,
			Method:      params.method,
			QueryParams: params.queryParams,
			Headers:     params.headers,
			Body:        params.body,
		},
		Response: domain.Response{
			Status: 200,
			Headers: domain.ResponseHeaders{
				"header1": "value1",
				"header2": "value2",
			},
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
						"stringKey": "arrayObjectStringValue1",
					},
					map[string]any{
						"stringKey": "arrayObjectStringValue2",
					},
				},
			},
		},
	}
}

func TestToSpec(t *testing.T) {
	sut := domain.Contracts{contract, contract}

	actual := sut.ToSpec("Consumer API", "Provider API")

	assert.Equal(t, domain.SpecHeading("Consumer APIが依存しているProvider APIの仕様"), actual.Heading)
	assert.Len(t, actual.Scenarios, 2)
}

func TestToScenario_シナリオ名にurlPathを使う(t *testing.T) {
	sut := createContract(SutParams{
		method:  "GET",
		urlPath: "/test",
		queryParams: domain.QueryParams{
			"query": {
				"equalTo": "value",
			},
		},
	})

	actual := sut.ToScenario("Consumer API")

	assert.Equal(t, domain.ScenarioHeading(`GET /test?query=value`), actual.Heading)
}

func TestToScenario_シナリオ名にurlを使う(t *testing.T) {
	sut := createContract(SutParams{
		method: "GET",
		url:    "/test",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Equal(t, domain.ScenarioHeading(`GET /test`), actual.Heading)
}

func TestToScenario_リクエストボディがある場合シナリオ名にボディのハッシュを含む(t *testing.T) {
	sut := createContract(SutParams{
		method: "PUT",
		url:    "/test",
		body:   "body",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Equal(t, domain.ScenarioHeading(`PUT /test (body: 841a2d689ad86bd1611447453c22c6fc)`), actual.Heading)
}

func TestToScenario_リクエストパスにurlPathを使う(t *testing.T) {
	sut := createContract(SutParams{
		method:  "GET",
		urlPath: "/test",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にGETリクエストを送る`))
}

func TestToScenario_リクエストパスにurlを使う(t *testing.T) {
	sut := createContract(SutParams{
		method: "GET",
		url:    "/test",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にGETリクエストを送る`))
}

func TestToScenario_GETリクエスト(t *testing.T) {
	sut := createContract(SutParams{
		method:  "GET",
		urlPath: "/test",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にGETリクエストを送る`))
}

func TestToScenario_POSTリクエスト(t *testing.T) {
	sut := createContract(SutParams{
		method:  "POST",
		urlPath: "/test",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にPOSTリクエストを送る`))
}

func TestToScenario_PUTリクエスト(t *testing.T) {
	sut := createContract(SutParams{
		method:  "PUT",
		urlPath: "/test",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にPUTリクエストを送る`))
}

func TestToScenario_DELETEリクエスト(t *testing.T) {
	sut := createContract(SutParams{
		method:  "DELETE",
		urlPath: "/test",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にDELETEリクエストを送る`))
}

func TestToScenario_クエリパラメータを含むリクエスト(t *testing.T) {
	sut := createContract(SutParams{
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

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test?q1=v1&q2=v2"にGETリクエストを送る`))
}

func TestToScenario_ヘッダを含むリクエスト(t *testing.T) {
	sut := createContract(SutParams{
		method:  "PUT",
		urlPath: "/test",
		headers: domain.RequestHeaders{
			"content-type": {
				"equalTo": "application/json",
			},
			"options": {
				"equalTo": "123, 456",
			},
		},
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にヘッダー"content-type: application/json \r\n options: 123, 456"で、PUTリクエストを送る`))
}

func TestToScenario_ボディを含むPOSTリクエスト(t *testing.T) {
	sut := createContract(SutParams{
		method:  "POST",
		urlPath: "/test",
		body:    "body",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(
		`URL"/test"にボディ"<file:fixtures/contracts/consumer_api/post_test_841a2d689ad86bd1611447453c22c6fc.json>"で、POSTリクエストを送る`))
}

func TestToScenario_ボディを含むPUTリクエスト(t *testing.T) {
	sut := createContract(SutParams{
		method:  "PUT",
		urlPath: "/test",
		body:    "body",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(
		`URL"/test"にボディ"<file:fixtures/contracts/consumer_api/put_test_841a2d689ad86bd1611447453c22c6fc.json>"で、PUTリクエストを送る`))
}

func TestToScenario_POSTとPUTリクエスト以外にはボディを含められない(t *testing.T) {
	sut := createContract(SutParams{
		method:  "GET",
		urlPath: "/test",
		body:    "body",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`URL"/test"にGETリクエストを送る`))
}

func TestToScenario_ヘッダとボディを含むPOSTリクエスト(t *testing.T) {
	sut := createContract(SutParams{
		method:  "POST",
		urlPath: "/test",
		headers: domain.RequestHeaders{
			"content-type": {
				"equalTo": "application/json",
			},
		},
		body: "body",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(
		`URL"/test"にボディ"<file:fixtures/contracts/consumer_api/post_test_841a2d689ad86bd1611447453c22c6fc.json>"、ヘッダー"content-type: application/json"で、POSTリクエストを送る`))
}

func TestToScenario_ヘッダとボディを含むPUTリクエスト(t *testing.T) {
	sut := createContract(SutParams{
		method:  "PUT",
		urlPath: "/test",
		headers: domain.RequestHeaders{
			"content-type": {
				"equalTo": "application/json",
			},
		},
		body: "body",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(
		`URL"/test"にボディ"<file:fixtures/contracts/consumer_api/put_test_841a2d689ad86bd1611447453c22c6fc.json>"、ヘッダー"content-type: application/json"で、PUTリクエストを送る`))
}

func TestToScenario_ボディを含む複雑なURLのPUTリクエスト(t *testing.T) {
	sut := createContract(SutParams{
		method: "PUT",
		url:    "/v1/companies/12345678?lang=ja&currency=JPY",
		body:   "body",
	})

	actual := sut.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(
		`URL"/v1/companies/12345678?lang=ja&currency=JPY"にボディ"<file:fixtures/contracts/consumer_api/put_v1_companies_12345678_lang_ja_currency_JPY_841a2d689ad86bd1611447453c22c6fc.json>"で、PUTリクエストを送る`))
}

func TestToScenario_レスポンスステータスコード(t *testing.T) {
	actual := contract.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスステータスコードが"200"である`))
}

func TestToScenario_レスポンスヘッダー(t *testing.T) {
	actual := contract.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスヘッダーに"header1"が存在し、その値が"value1"である`))
	assert.Contains(t, actual.Steps, domain.Step(`レスポンスヘッダーに"header2"が存在し、その値が"value2"である`))
}

func TestToScenario_文字列のアサーション(t *testing.T) {
	actual := contract.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.stringKey"が文字列の"stringValue"である`))
}

func TestToScenario_整数のアサーション(t *testing.T) {
	actual := contract.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.integerKey"が整数の"123"である`))
}

func TestToScenario_小数のアサーション(t *testing.T) {
	actual := contract.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.floatKey"が小数の"123.456"である`))
}

func TestToScenario_真偽値のアサーション(t *testing.T) {
	actual := contract.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.booleanKey"が真偽値の"true"である`))
}

func TestToScenario_オブジェクトに含まれる値のアサーション(t *testing.T) {
	actual := contract.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.objectKey.stringKey"が文字列の"objectStringValue"である`))
}

func TestToScenario_配列に含まれるオブジェクトに含まれる値のアサーション(t *testing.T) {
	actual := contract.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.arrayKey[0].stringKey"が文字列の"arrayObjectStringValue1"である`))
}

func TestToScenario_配列に含まれる2つめ以降のオブジェクトに含まれる値のアサーション(t *testing.T) {
	actual := contract.ToScenario("Consumer API")

	assert.Contains(t, actual.Steps, domain.Step(`レスポンスのJSONの"$.arrayKey[1].stringKey"が文字列の"arrayObjectStringValue2"である`))
}

func TestToScenario_レスポンスボディのアサーションはキーの昇順で並べる(t *testing.T) {
	sut := &domain.Contract{
		Request: domain.Request{
			Url:    "/test",
			Method: "GET",
		},
		Response: domain.Response{
			Status: 200,
			JsonBody: map[string]any{
				"c": "c value",
				"b": "b value",
				"a": map[string]any{
					"x": "a.x value",
				},
				"abc": "abc value",
			},
		},
	}

	actual := sut.ToScenario("Consumer API")

	// Steps[0]はリクエスト、Steps[1]はステータスコード
	assert.Contains(t, actual.Steps[2], "a.x value")
	assert.Contains(t, actual.Steps[3], "abc value")
	assert.Contains(t, actual.Steps[4], "b value")
	assert.Contains(t, actual.Steps[5], "c value")
}
