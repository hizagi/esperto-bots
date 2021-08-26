package utils

import (
	"os"
	"path/filepath"
)

func FromRootPath(toRootPath string, goToDir string) string {
	currentPath, _ := os.Getwd()

	rootPath := (filepath.ToSlash(currentPath) + "/" + toRootPath)

	return rootPath + goToDir
}
