package steamworkshop

import (
	"fmt"

	"github.com/CarlFlo/steamWorkshopDownloader/database"
)

var (
	errFailedToFetchAppID  = fmt.Errorf("failed to fetch the appID")
	errMissingItemName     = fmt.Errorf("failed to fetch item name")
	errMissingLastUpdated  = fmt.Errorf("failed to fetch last updated")
	errMissingCreatedBy    = fmt.Errorf("failed to fetch created by")
	errMissingCreatorName  = fmt.Errorf("failed to fetch creator name")
	errMissingFileSize     = fmt.Errorf("failed to fetch file size")
	errMissingPreviewImage = fmt.Errorf("failed to fetch preview image URL")
)

func ParseWorkshopURL(url, workshopID string) (*database.WorkshopItem, error) {

	workshopData, err := visitPage(url)
	if err != nil {
		return nil, err
	}

	workshopData.WorkshopID = workshopID

	return workshopData, nil
}

func visitPage(url string) (*database.WorkshopItem, error) {

	info, err := getWorkshopPageInfo(url)
	if err != nil {
		return nil, err
	}

	// validate
	if len(info.AppID) == 0 {
		return nil, errFailedToFetchAppID
	} else if len(info.ItemName) == 0 {
		return nil, errMissingItemName
	} else if len(info.LastUpdated) == 0 {
		return nil, errMissingLastUpdated
	} else if len(info.CreatorLink) == 0 {
		return nil, errMissingCreatedBy
	} else if len(info.CreatorName) == 0 {
		return nil, errMissingCreatorName
	} else if len(info.FileSize) == 0 {
		return nil, errMissingFileSize
	} else if len(info.PreviewImage) == 0 {
		return nil, errMissingPreviewImage
	}

	return info, err
}
