package domain_test

import (
	"play-cdc/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUrl_urlがある場合はurlを使う(t *testing.T) {
	actual := domain.NewUrl(
		domain.StubRequest{Url: "/stub"},
		domain.ActualRequest{Url: "/actual"},
	)

	assert.Equal(t, "/stub", actual.String())
}

func TestNewUrl_urlがなくurlPathがある場合はurlPathを使う(t *testing.T) {
	actual := domain.NewUrl(
		domain.StubRequest{UrlPath: "/stub/path"},
		domain.ActualRequest{Url: "/actual"},
	)

	assert.Equal(t, "/stub/path", actual.String())
}

func TestNewUrl_urlPathとクエリパラメータを結合する(t *testing.T) {
	actual := domain.NewUrl(
		domain.StubRequest{
			UrlPath: "/test",
			QueryParams: domain.QueryParams{
				"q1": {"equalTo": "v1"},
				"q2": {"equalTo": "v2"},
			},
		},
		domain.ActualRequest{},
	)

	assert.Equal(t, "/test?q1=v1&q2=v2", actual.String())
}

func TestNewUrl_urlもurlPathもない場合は実際のリクエストURLを使う(t *testing.T) {
	actual := domain.NewUrl(
		domain.StubRequest{},
		domain.ActualRequest{Url: "/users/123?active=true"},
	)

	assert.Equal(t, "/users/123?active=true", actual.String())
}

func TestNewUrl_urlがある場合は実際のリクエストURLで上書きしない(t *testing.T) {
	actual := domain.NewUrl(
		domain.StubRequest{Url: "/users/stub"},
		domain.ActualRequest{Url: "/users/123"},
	)

	assert.Equal(t, "/users/stub", actual.String())
}

func TestNewUrl_urlPathがある場合は実際のリクエストURLで上書きしない(t *testing.T) {
	actual := domain.NewUrl(
		domain.StubRequest{UrlPath: "/users/stub"},
		domain.ActualRequest{Url: "/users/123"},
	)

	assert.Equal(t, "/users/stub", actual.String())
}

func TestNewUrl_urlとurlPathの両方がある場合はurlを優先する(t *testing.T) {
	actual := domain.NewUrl(
		domain.StubRequest{
			Url:     "/url1",
			UrlPath: "/url1/path",
		},
		domain.ActualRequest{},
	)

	assert.Equal(t, "/url1", actual.String())
}
