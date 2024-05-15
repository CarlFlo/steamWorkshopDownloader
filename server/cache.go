package server

import (
	"os"
)

// Returns true if the file exists in the cache
func lookForFileOnDisk(fi *FileInfo) bool {

	return fileExists(fi.ZipFilePath)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
