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
	v1, v2 := "v1", "v2"
	actual := domain.NewUrl(
		domain.StubRequest{
			UrlPath: "/test",
			QueryParams: domain.QueryParams{
				"q1": {EqualTo: &v1},
				"q2": {EqualTo: &v2},
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

func TestNewUrl_equalTo以外のクエリ条件が混在する場合は実際のリクエストURLを使う(t *testing.T) {
	v := "v1"
	actual := domain.NewUrl(
		domain.StubRequest{
			UrlPath: "/test",
			QueryParams: domain.QueryParams{
				"q1": {EqualTo: &v},
				"q2": {},
			},
		},
		domain.ActualRequest{Url: "/test?q1=v1&q2=actual&extra=1"},
	)

	assert.Equal(t, "/test?q1=v1&q2=actual&extra=1", actual.String())
}

func TestNewUrl_equalToが一つもないクエリ条件は実際のリクエストURLを使う(t *testing.T) {
	actual := domain.NewUrl(
		domain.StubRequest{
			UrlPath: "/test",
			QueryParams: domain.QueryParams{
				"q": {},
			},
		},
		domain.ActualRequest{Url: "/test?q=actual"},
	)

	assert.Equal(t, "/test?q=actual", actual.String())
}

func TestNewUrl_urlとurlPathの両方がある場合はurlを優先する(t *testing.T) {
	actual := domain.NewUrl(
		domain.StubRequest{
			Url:     "/url1",
			UrlPath: "/url1/path",
			QueryParams: domain.QueryParams{
				"q": {},
			},
		},
		domain.ActualRequest{Url: "/actual?q=value"},
	)

	assert.Equal(t, "/url1", actual.String())
}
