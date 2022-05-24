package domain_test

import (
	"specify/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSpec(t *testing.T) {
	contract := domain.Contract{
		domain.Request{
			"/test",
			"GET",
		},
		domain.Response{
			200,
			nil,
		},
	}
	actual := domain.GenerateSpec(contract)

	expected := `## GET /test
* URL"/test"にGETリクエストを送る
* レスポンスステータスコードが"200"である
`
	assert.Equal(t, expected, actual)
}
