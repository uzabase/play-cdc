package domain_test

import (
	"errors"
	"play-cdc/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinErrors(t *testing.T) {
	tests := []struct {
		name   string
		errors []error
		want   string
	}{
		{
			name:   "空のスライス",
			errors: []error{},
			want:   "",
		},
		{
			name:   "1件のエラー",
			errors: []error{errors.New("error1")},
			want:   "error1",
		},
		{
			name:   "複数のエラー",
			errors: []error{errors.New("error1"), errors.New("error2"), errors.New("error3")},
			want:   "error1; error2; error3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, domain.JoinErrors(tt.errors))
		})
	}
}
