package server

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

func secureFileServer(w http.ResponseWriter, r *http.Request, basePath, relPath string) {
	path := filepath.Join(basePath, relPath)

	err := checkPath(basePath, path)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.ServeFile(w, r, path)
}

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
