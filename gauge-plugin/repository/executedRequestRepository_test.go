package repository_test

import (
	"fmt"
	"play-cdc/domain"
	"play-cdc/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshal_Jsonのレスポンスボディ(t *testing.T) {
	json := `{
		"requests" : [ {
		  "stubMapping" : {
			"response" : {
			  "jsonBody" : {
				"globalId": "JP3039710003"
			  }
			}
		  }
		} ]
	  }`

	result, err := repository.Unmarshal([]byte(json))

	expected := &domain.ExecutedRequests{
		Requests: []domain.ExecutedRequest{
			{
				StubMapping: domain.StubMapping{
					Response: domain.StubResponse{
						JsonBody: map[string]any{
							"globalId": "JP3039710003",
						},
					},
				},
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUnmarshal_ルート要素が配列になっているJsonのレスポンスボディ(t *testing.T) {
	json := `{
		"requests" : [ {
		  "stubMapping" : {
			"response" : {
			  "jsonBody" : [
				{
				  "globalId": "JP3039710003"
			    }
			  ]
			}
		  }
		} ]
	  }`

	result, err := repository.Unmarshal([]byte(json))

	expected := &domain.ExecutedRequests{
		Requests: []domain.ExecutedRequest{
			{
				StubMapping: domain.StubMapping{
					Response: domain.StubResponse{
						JsonBody: []any{
							map[string]any{
								"globalId": "JP3039710003",
							},
						},
					},
				},
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUnmarshal_空文字列の場合はエラー(t *testing.T) {
	json := ""

	result, err := repository.Unmarshal([]byte(json))

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestUnmarshal_クエリ条件から契約のURLを生成する(t *testing.T) {
	tests := []struct {
		name            string
		queryParameters string
		actualURL       string
		expectedURL     string
	}{
		{
			name:            "文字列のtrueは値として使う",
			queryParameters: `{"q":{"equalTo":"true"}}`,
			actualURL:       "/test?q=true&extra=value",
			expectedURL:     "/test?q=true",
		},
		{
			name:            "空文字列はパラメータとして残す",
			queryParameters: `{"q":{"equalTo":""}}`,
			actualURL:       "/test?q=",
			expectedURL:     "/test?q=",
		},
		{
			name:            "absentのパラメータは疑問符も含めて省略する",
			queryParameters: `{"q":{"absent":true}}`,
			actualURL:       "/test",
			expectedURL:     "/test",
		},
		{
			name:            "複数のabsentでも疑問符を付けない",
			queryParameters: `{"q1":{"absent":true},"q2":{"absent":true}}`,
			actualURL:       "/test",
			expectedURL:     "/test",
		},
		{
			name:            "absentと通常の条件を混在させる",
			queryParameters: `{"q2":{"equalTo":"v2"},"omitted":{"absent":true},"q1":{"equalTo":"v1"}}`,
			actualURL:       "/test?q2=v2&q1=v1",
			expectedURL:     "/test?q1=v1&q2=v2",
		},
		{
			name:            "caseInsensitiveの真偽値を読み込める",
			queryParameters: `{"q":{"equalTo":"TRUE","caseInsensitive":true}}`,
			actualURL:       "/test?q=true",
			expectedURL:     "/test?q=TRUE",
		},
		{
			name:            "未知の真偽値や配列やオブジェクトのオプションを読み飛ばす",
			queryParameters: `{"q":{"equalTo":"value","futureFlag":true,"futureArray":[true,{"enabled":true}],"futureObject":{"enabled":true}}}`,
			actualURL:       "/test?q=value",
			expectedURL:     "/test?q=value",
		},
		{
			name:            "falseのオプションも読み込める",
			queryParameters: `{"q":{"equalTo":"true","caseInsensitive":false,"absent":false}}`,
			actualURL:       "/test?q=true",
			expectedURL:     "/test?q=true",
		},
		{
			name:            "値を指定しないabsent_falseは実際のURLを使う",
			queryParameters: `{"q":{"absent":false}}`,
			actualURL:       "/test?q=value",
			expectedURL:     "/test?q=value",
		},
		{
			name:            "正規表現は実際のURLを使いエンコードと複数値を保持する",
			queryParameters: `{"q":{"matches":".+"}}`,
			actualURL:       "/test?q=a%26b&q=c%20d",
			expectedURL:     "/test?q=a%26b&q=c%20d",
		},
		{
			name:            "値を決められない条件が混在する場合はURL全体を実際のURLから取る",
			queryParameters: `{"q":{"matches":".+"},"exact":{"equalTo":"value"},"omitted":{"absent":true}}`,
			actualURL:       "/test?q=actual&exact=value",
			expectedURL:     "/test?q=actual&exact=value",
		},
		{
			name:            "配列を含む複合条件も実際のURLを使う",
			queryParameters: `{"q":{"or":[{"equalTo":"value"},{"absent":true}]}}`,
			actualURL:       "/test?q=value",
			expectedURL:     "/test?q=value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			json := fmt.Sprintf(`{
				"requests": [{
					"request": {"url": %q},
					"wasMatched": true,
					"stubMapping": {
						"request": {
							"method": "GET",
							"urlPath": "/test",
							"queryParameters": %s
						},
						"response": {"status": 200}
					}
				}]
			}`, tt.actualURL, tt.queryParameters)

			requests, err := repository.Unmarshal([]byte(json))
			require.NoError(t, err)

			contracts := domain.ToContracts(requests.Requests)
			require.Len(t, contracts, 1)
			assert.Equal(t, tt.expectedURL, contracts[0].Request.Url.String())
			assert.Equal(t, fmt.Sprintf(
				"# consumerが依存しているproviderの仕様\n\n## GET %s\n* URL\"%s\"にGETリクエストを送る\n* レスポンスステータスコードが\"200\"である\n",
				tt.expectedURL, tt.expectedURL,
			), contracts.ToSpec("consumer", "provider").String())
		})
	}
}

func TestUnmarshal_equalToの値が真偽値の場合はエラー(t *testing.T) {
	json := `{
		"requests": [{
			"stubMapping": {
				"request": {"queryParameters": {"q": {"equalTo": true}}}
			}
		}]
	}`

	result, err := repository.Unmarshal([]byte(json))

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUnmarshal_実際のURLにフォールバックした契約をURLごとに区別する(t *testing.T) {
	json := `{
		"requests": [
			{
				"request": {"url": "/test?q=first"},
				"wasMatched": true,
				"stubMapping": {
					"request": {"method": "GET", "urlPath": "/test", "queryParameters": {"q": {"matches": ".+"}}},
					"response": {"status": 200}
				}
			},
			{
				"request": {"url": "/test?q=second"},
				"wasMatched": true,
				"stubMapping": {
					"request": {"method": "GET", "urlPath": "/test", "queryParameters": {"q": {"matches": ".+"}}},
					"response": {"status": 200}
				}
			},
			{
				"request": {"url": "/test?q=first"},
				"wasMatched": true,
				"stubMapping": {
					"request": {"method": "GET", "urlPath": "/test", "queryParameters": {"q": {"matches": ".+"}}},
					"response": {"status": 200}
				}
			}
		]
	}`

	requests, err := repository.Unmarshal([]byte(json))
	require.NoError(t, err)

	contracts := domain.ToContracts(requests.Requests)
	require.Len(t, contracts, 2)
	assert.Equal(t, "/test?q=first", contracts[0].Request.Url.String())
	assert.Equal(t, "/test?q=second", contracts[1].Request.Url.String())
}
