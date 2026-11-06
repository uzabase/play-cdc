package repository_test

import (
	"play-cdc/domain"
	"play-cdc/repository"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestUnmarshal_クエリ条件のequalToを読み込む(t *testing.T) {
	json := `{
		"requests": [{
			"stubMapping": {
				"request": {
					"queryParameters": {
						"q": {"equalTo": "v1"}
					}
				}
			}
		}]
	}`

	result, err := repository.Unmarshal([]byte(json))

	q := "v1"
	expected := &domain.ExecutedRequests{
		Requests: []domain.ExecutedRequest{
			{
				StubMapping: domain.StubMapping{
					Request: domain.StubRequest{
						QueryParams: domain.QueryParams{
							"q": {EqualTo: &q},
						},
					},
				},
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUnmarshal_equalToがないクエリ条件を読み飛ばす(t *testing.T) {
	json := `{
		"requests": [{
			"stubMapping": {
				"request": {
					"queryParameters": {
						"debug": {"absent": true}
					}
				}
			}
		}]
	}`

	result, err := repository.Unmarshal([]byte(json))

	expected := &domain.ExecutedRequests{
		Requests: []domain.ExecutedRequest{
			{
				StubMapping: domain.StubMapping{
					Request: domain.StubRequest{
						QueryParams: domain.QueryParams{
							"debug": {},
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
