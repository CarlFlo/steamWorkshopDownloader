package server

import (
	"os"

	steamworkshop "github.com/CarlFlo/steamWorkshopDownloader/steamWorkshop"
)

// Returns true if the file exists in the cache
func checkCache(fi *FileInfo) bool {

	return fileExists(fi.ZipFilePath)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func LoadWorkshopData(itemID string) *steamworkshop.WorkshopData {

	// Check database
	return &steamworkshop.WorkshopData{}
}
