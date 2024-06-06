package domain_test

import (
	"play-cdc/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToContracts(t *testing.T) {
	sut := []domain.ExecutedRequest{
		{
			Request: domain.ActualRequest{
				Body: "body",
			},
			WasMatched: true,
			StubMapping: domain.StubMapping{
				Request: domain.StubRequest{
					Method:      "GET",
					Url:         "/url1",
					UrlPath:     "/url1/path",
					QueryParams: domain.QueryParams{},
					Headers:     domain.RequestHeaders{},
				},
				Response: domain.StubResponse{
					Status: 200,
					Headers: domain.ResponseHeaders{
						"headerKey": "headerValue",
					},
					JsonBody: map[string]any{
						"bodyKey": "bodyValue",
					},
				},
			},
		},
	}

	actual := domain.ToContracts(sut)

	expected := domain.Contracts{
		&domain.Contract{
			Request: domain.Request{
				Method:      "GET",
				Url:         "/url1",
				UrlPath:     "/url1/path",
				QueryParams: domain.QueryParams{},
				Headers:     domain.RequestHeaders{},
				Body:        "body",
			},
			Response: domain.Response{
				Status: 200,
				Headers: domain.ResponseHeaders{
					"headerKey": "headerValue",
				},
				Body: domain.NewJsonResponseBody(
					map[string]any{
						"bodyKey": "bodyValue",
					},
				),
			},
		},
	}
	assert.Equal(t, expected, actual)
}

func TestToContracts_JsonBodyのルート要素が配列(t *testing.T) {
	sut := []domain.ExecutedRequest{
		{
			WasMatched: true,
			StubMapping: domain.StubMapping{
				Response: domain.StubResponse{
					JsonBody: []any{
						map[string]any{
							"bodyKey": "bodyValue",
						},
					},
				},
			},
		},
	}
	expected := domain.Contracts{
		&domain.Contract{
			Response: domain.Response{
				Body: domain.NewJsonResponseBody(
					[]any{
						map[string]any{
							"bodyKey": "bodyValue",
						},
					},
				),
			},
		},
	}
	actual := domain.ToContracts(sut)

	assert.Equal(t, expected, actual)
}

func TestToContracts_JsonBodyがなくBodyがJsonの場合パースして使う(t *testing.T) {
	sut := []domain.ExecutedRequest{
		{
			WasMatched: true,
			StubMapping: domain.StubMapping{
				Response: domain.StubResponse{
					Body: "{\n  \"key\": \"value\"\n}",
				},
			},
		},
	}
	expected := domain.Contracts{
		&domain.Contract{
			Response: domain.Response{
				Body: domain.NewJsonResponseBody(
					map[string]any{
						"key": "value",
					},
				),
			},
		},
	}
	actual := domain.ToContracts(sut)

	assert.Equal(t, expected, actual)
}

func TestToContracts_JsonBodyがなくBodyがJsonではない場合テキストとして扱う(t *testing.T) {
	sut := []domain.ExecutedRequest{
		{
			WasMatched: true,
			StubMapping: domain.StubMapping{
				Response: domain.StubResponse{
					Body: "some text message",
				},
			},
		},
	}
	expected := domain.Contracts{
		&domain.Contract{
			Response: domain.Response{
				Body: domain.NewTextResponseBody("some text message"),
			},
		},
	}
	actual := domain.ToContracts(sut)

	assert.Equal(t, expected, actual)
}

func TestToContracts_JsonBodyがない場合のBodyのルート要素が配列(t *testing.T) {
	sut := []domain.ExecutedRequest{
		{
			WasMatched: true,
			StubMapping: domain.StubMapping{
				Response: domain.StubResponse{
					Body: "[\n  {\n    \"key\": \"value\"\n  }\n]",
				},
			},
		},
	}
	expected := domain.Contracts{
		&domain.Contract{
			Response: domain.Response{
				Body: domain.NewJsonResponseBody(
					[]any{
						map[string]any{
							"key": "value",
						},
					},
				),
			},
		},
	}
	actual := domain.ToContracts(sut)

	assert.Equal(t, expected, actual)
}

func TestToContracts_マッチしたリクエストだけを契約として扱う(t *testing.T) {
	sut := []domain.ExecutedRequest{
		{
			WasMatched: true,
			StubMapping: domain.StubMapping{
				Request: domain.StubRequest{
					Url: "/url1",
				},
			},
		},
		{
			WasMatched: false,
			StubMapping: domain.StubMapping{
				Request: domain.StubRequest{
					Url: "/url2",
				},
			},
		},
	}

	actual := domain.ToContracts(sut)

	expected := domain.Contracts{
		&domain.Contract{
			Request: domain.Request{
				Url: "/url1",
			},
		},
	}
	assert.Equal(t, expected, actual)
}

func TestToContracts_契約は契約全体でユニークにする(t *testing.T) {
	sut := []domain.ExecutedRequest{
		{
			WasMatched: true,
			StubMapping: domain.StubMapping{
				Request: domain.StubRequest{
					Method: "GET",
				},
			},
		},
		{
			WasMatched: true,
			StubMapping: domain.StubMapping{
				Request: domain.StubRequest{
					Method: "PUT",
				},
			},
		},
		{
			WasMatched: true,
			StubMapping: domain.StubMapping{
				Request: domain.StubRequest{
					Method: "GET",
				},
			},
		},
	}

	actual := domain.ToContracts(sut)

	assert.Len(t, actual, 2)
}
