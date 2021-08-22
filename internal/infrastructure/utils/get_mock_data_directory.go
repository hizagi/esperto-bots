package utils

import (
	"os"
	"path/filepath"
)

const mockDir = "seeds"

func GetMockDataDirectory(rootPath string, mockDataFilename string) string {
	path, _ := os.Getwd()
	path = (filepath.ToSlash(path) + "/" + rootPath + "/" + mockDir + "/" + mockDataFilename)

	return path
}
