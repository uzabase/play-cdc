package repository

import (
	"fmt"
	"os"
	"play-cdc/domain"
)

func SaveSpec(spec *domain.Spec, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create a file for spec(%s): %w", spec.Heading, err)
	}
	defer f.Close()

	_, err = f.WriteString(spec.String())
	if err != nil {
		return fmt.Errorf("Failed to write to a file for spec(%s): %w", spec.Heading, err)
	}

	return nil
}
