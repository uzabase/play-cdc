package repository

import (
	"fmt"
	"specify/domain"
)

func SaveSpec(spec *domain.Spec) error {
	fmt.Println(spec)
	return nil
}
