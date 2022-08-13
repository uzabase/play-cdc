package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"play-cdc/domain"
)

func SaveSpec(spec *domain.Spec, path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return fmt.Errorf("Error: failed to create a directory(%s): %w", dir, err)
	}

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
