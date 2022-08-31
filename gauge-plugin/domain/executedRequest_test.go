package domain_test

import (
	"play-cdc/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToContracts(t *testing.T) {
	sut := domain.ExecutedRequests{
		Requests: []domain.ExecutedRequest{
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
				},
			},
		},
	}

	actual := sut.ToContracts()

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
		},
	}
	assert.Equal(t, expected, actual)
}

func TestToContracts_マッチしたリクエストだけを契約として扱う(t *testing.T) {
	sut := domain.ExecutedRequests{
		Requests: []domain.ExecutedRequest{
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
		},
	}

	actual := sut.ToContracts()

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
	sut := domain.ExecutedRequests{
		Requests: []domain.ExecutedRequest{
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
		},
	}

	actual := sut.ToContracts()

	assert.Len(t, actual, 2)
}
