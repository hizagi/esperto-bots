package utils

import (
	"os"
	"path/filepath"
)

const mockDir = "seeds"
const databaseContainerSeedDir = "/docker-entrypoint-initdb.d/"

func GetMockDataDirectory(rootPath string, mockFiles []string) map[string]string {
	paths := make(map[string]string, 0)

	for _, mockFile := range mockFiles {

		currentPath, _ := os.Getwd()
		path := (filepath.ToSlash(currentPath) + "/" + rootPath + "/" + mockDir + "/" + mockFile)

		paths[path] = databaseContainerSeedDir + mockFile
	}

	return paths
}
