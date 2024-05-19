package utils

import (
	"os"
)

// Returns true if the file exists in the cache
func LookForFileOnDisk(path string) bool {

	return fileExists(path)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
