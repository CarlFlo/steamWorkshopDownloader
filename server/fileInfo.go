package server

import (
	"fmt"
	"strings"

	"github.com/CarlFlo/steamWorkshopDownloader/config"
)

type FileInfo struct {
	Url         string
	WorkshopID  string
	ZipFileName string
	ZipFilePath string
}

// New - Creates a new struct. The URL argument is the URL to the workshop item. Example: https://steamcommunity.com/sharedfiles/filedetails/?id=...
func (fi *FileInfo) New(url string) {

	fi.Url = url

	// Get workshop ID from the URL
	split := strings.Split(fi.Url, "=")
	fi.WorkshopID = split[len(split)-1]

	// Name of the .zip file - format <ID>.zip
	fi.ZipFileName = fmt.Sprintf("%s.zip", fi.WorkshopID)

	// Filepath to the .ZIP file
	fi.ZipFilePath = fmt.Sprintf("%s\\%s", config.CONFIG.PathToCache, fi.ZipFileName)
}

func (fi *FileInfo) ToString() string {
	return fmt.Sprintf("url: %s\nWorkshopID: %s\nZipFileName: %s\nZipFilePath: %s\n", fi.Url, fi.WorkshopID, fi.ZipFileName, fi.ZipFilePath)
}
