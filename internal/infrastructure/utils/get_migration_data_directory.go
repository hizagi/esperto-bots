package utils

import (
	"os"

	"github.com/hizagi/esperto-bots/projectpath"
)

const databaseContainerMigrationDir = "/docker-entrypoint-initdb.d/"

func GetMigrationsDataDirectory(databaseName string) map[string]string {
	paths := make(map[string]string, 0)
	basePath := projectpath.Root + "/internal/infrastructure/database/" + databaseName + "/migrations/"

	migrationFiles, err := oSReadDir(basePath)

	if err != nil {
		panic(err)
	}

	for _, migrationFile := range migrationFiles {

		path := basePath + migrationFile

		paths[path] = databaseContainerSeedDir + migrationFile
	}

	return paths
}

func oSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
