package utils

import "os"

func GetFilenamesFromDir(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return []string{}
	}
	defer file.Close()
	names, err := file.Readdirnames(0)
	if err != nil {
		return []string{}
	}

	return names
}
