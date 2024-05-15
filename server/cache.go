package server

import (
	"os"

	"github.com/CarlFlo/steamWorkshopDownloader/database"
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

func LoadWorkshopData(itemID string) *database.WorkshopItem {

	// Check database
	return &database.WorkshopItem{}
}
