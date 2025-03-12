package server

import (
	"fmt"
	"path/filepath"
	"strings"
)

func safeJoin(basePath, path string) (string, error) {
	joinedPath := filepath.Join(basePath, path)

	absBasePath, err := filepath.Abs(basePath)
	if err != nil {
		return "", err
	}

	absJoinedPath, err := filepath.Abs(joinedPath)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(absJoinedPath, absBasePath) {
		return "", fmt.Errorf("user tries to escape the base directory (%s)", path)
	}

	return joinedPath, nil
}
