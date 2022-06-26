package domain_test

import (
	"specify/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToContracts_マッチしたリクエストだけを契約として扱う(t *testing.T) {
	sut := domain.ExecutedRequests{
		Requests: []domain.ExecutedRequest{
			{
				WasMatched: true,
				StubMapping: domain.StubMapping{
					Id: "id1",
					Request: domain.Request{
						Url: "/url1",
					},
				},
			},
			{
				WasMatched: false,
				StubMapping: domain.StubMapping{
					Id: "id2",
					Request: domain.Request{
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

func TestToContracts_契約はStubMappingのIdでユニークにする(t *testing.T) {
	sut := domain.ExecutedRequests{
		Requests: []domain.ExecutedRequest{
			{
				WasMatched: true,
				StubMapping: domain.StubMapping{
					Id: "id1",
				},
			},
			{
				WasMatched: true,
				StubMapping: domain.StubMapping{
					Id: "id1",
				},
			},
		},
	}

	actual := sut.ToContracts()

	assert.Len(t, actual, 1)
}
