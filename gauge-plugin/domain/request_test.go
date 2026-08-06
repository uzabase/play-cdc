package domain_test

import (
	"play-cdc/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBodyFileName(t *testing.T) {
	sut := domain.Request{
		Method: "POST",
		Url:    "/users/123",
		Body:   "body",
	}

	actual := sut.ToBodyFileName()

	assert.Equal(t, "post_users_123_841a2d689ad86bd1611447453c22c6fc.json", actual)
}

func TestToBodyFileName_URLが空でもpanicしない(t *testing.T) {
	sut := domain.Request{
		Method: "POST",
		Body:   "body",
	}

	assert.NotPanics(t, func() {
		actual := sut.ToBodyFileName()
		assert.Equal(t, "post__841a2d689ad86bd1611447453c22c6fc.json", actual)
	})
}
