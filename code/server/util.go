package server

import (
	"fmt"
	"path/filepath"
	"strings"
)

func checkPath(basePath, path string) error {
	absBasePath, err := filepath.Abs(basePath)
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(absPath, absBasePath) {
		return fmt.Errorf("user tries to escape the base directory (%s)", path)
	}

	return nil
}
