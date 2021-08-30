package utils

import "github.com/hizagi/esperto-bots/projectpath"

const databaseContainerSeedDir = "/docker-entrypoint-initdb.d/"

func GetSeedDataDirectory(databaseName string, seedFiles []string) map[string]string {
	paths := make(map[string]string, 0)

	for _, seedFile := range seedFiles {

		path := projectpath.Root + "/internal/infrastructure/database/" + databaseName + "/seeds/" + seedFile

		paths[path] = databaseContainerSeedDir + seedFile
	}

	return paths
}
