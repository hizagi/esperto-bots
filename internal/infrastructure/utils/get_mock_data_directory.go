package utils

const databaseContainerSeedDir = "/docker-entrypoint-initdb.d/"

func GetSeedDataDirectory(toRootPath string, seedDir string, seedFiles []string) map[string]string {
	paths := make(map[string]string, 0)

	for _, seedFile := range seedFiles {

		path := FromRootPath(toRootPath, seedDir+"/"+seedFile)

		paths[path] = databaseContainerSeedDir + seedFile
	}

	return paths
}
