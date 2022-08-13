package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"play-cdc/domain"
)

func SaveRequestBodies(contracts domain.Contracts, path string) {
	for _, c := range contracts {
		if !c.Request.IsBodyAvailable() {
			continue
		}

		if err := writeToFile(path, c.Request); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func writeToFile(path string, request domain.Request) error {
	if err := os.MkdirAll(path, 0777); err != nil {
		return fmt.Errorf("Error: failed to create a directory(%s): %w", path, err)
	}

	filePath := filepath.Join(path, request.ToBodyFileName())
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error: failed to create a file(%s): %w", filePath, err)
	}
	defer f.Close()

	if _, err = f.WriteString(request.Body); err != nil {
		return fmt.Errorf("Error: failed to write to a file(%s): %w", filePath, err)
	}

	return nil
}
