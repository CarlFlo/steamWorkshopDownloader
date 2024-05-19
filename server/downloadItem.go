package server

import (
	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/steamWorkshopDownloader/database"
	steamworkshop "github.com/CarlFlo/steamWorkshopDownloader/steamWorkshop"
	"github.com/CarlFlo/steamWorkshopDownloader/utils"
)

func prepareAndDownloadItem(fi *FileInfo) (*database.WorkshopItem, error) {

	// Look in the cache
	var workshopData database.WorkshopItem
	if err := workshopData.QueryItemByWorkshopID(fi.WorkshopID); err != nil {
		malm.Error("Could not fetch '%s' from database. Reason: %v", fi.WorkshopID, err)
	}

	item, err := steamworkshop.DownloadWorkshopItem(workshopData.GetCommand())
	if err != nil {
		return nil, err
	}

	// The zipFileName or ID is unique to every workshop file. Can be used for cache
	if err := utils.ZipFolder(item.PathToFile, fi.ZipFilePath); err != nil {
		return nil, err
	}

	return &workshopData, nil
}
