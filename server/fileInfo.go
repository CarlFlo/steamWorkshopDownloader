package server

import (
	"fmt"
	"net/url"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/steamWorkshopDownloader/config"
)

type FileInfo struct {
	Url         string
	WorkshopID  string
	ZipFileName string
	ZipFilePath string
}

// New - Creates a new struct. The URL argument is the URL to the workshop item. Example: https://steamcommunity.com/sharedfiles/filedetails/?id=...
func (fi *FileInfo) New(inputUrl string) {

	fi.Url = inputUrl

	// The url has already been checked by the "CheckUrlInput" function. There should not be any problems here

	parsedURL, err := url.Parse(fi.Url)
	if err != nil {
		malm.Error("failed to parse url. This should not have happened '%v'", err)
		return
	}

	queryParams := parsedURL.Query()
	id := queryParams.Get("id")

	if len(id) == 0 {
		malm.Error("failed to get id param from url. This should not have happened. %s", fi.Url)
		return
	}

	fi.WorkshopID = id

	// Name of the .zip file - format <ID>.zip
	fi.ZipFileName = fmt.Sprintf("%s.zip", fi.WorkshopID)

	// Filepath to the .ZIP file
	fi.ZipFilePath = fmt.Sprintf("%s\\%s", config.CONFIG.PathToCache, fi.ZipFileName)
}

func (fi *FileInfo) ToString() string {
	return fmt.Sprintf("url: %s\nWorkshopID: %s\nZipFileName: %s\nZipFilePath: %s\n", fi.Url, fi.WorkshopID, fi.ZipFileName, fi.ZipFilePath)
}
