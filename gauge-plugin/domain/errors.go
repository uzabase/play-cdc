package domain

import "strings"

func JoinErrors(errors []error) string {
	msgs := make([]string, len(errors))
	for i, err := range errors {
		msgs[i] = err.Error()
	}
	return strings.Join(msgs, "; ")
}
