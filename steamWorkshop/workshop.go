package steamworkshop

import (
	"fmt"

	"github.com/CarlFlo/malm"
)

type WorkshopData struct {
	AppID      string `json:"app_id"`
	WorkshopID string `json:"workshop_id"`

	ItemName     string `json:"item_name"`
	LastUpdated  string `json:"last_updated"`
	CreatedBy    string `json:"created_by"`
	CreatorName  string `json:"creator_name"`
	FileSize     string `json:"file_size"`
	PreviewImage string `json:"preview_image"`
}

func (item WorkshopData) GetCommand() string {

	return fmt.Sprintf("workshop_download_item %s %s", item.AppID, item.WorkshopID)
}

var (
	errFailedToFetchAppID  = fmt.Errorf("failed to fetch the appID")
	errMissingItemName     = fmt.Errorf("failed to fetch item name")
	errMissingLastUpdated  = fmt.Errorf("failed to fetch last updated")
	errMissingCreatedBy    = fmt.Errorf("failed to fetch created by")
	errMissingCreatorName  = fmt.Errorf("failed to fetch creator name")
	errMissingFileSize     = fmt.Errorf("failed to fetch file size")
	errMissingPreviewImage = fmt.Errorf("failed to fetch preview image URL")
)

func ParseWorkshopURL(url, workshopID string) (*WorkshopData, error) {

	workshopData, err := visitPage(url)
	if err != nil {
		return nil, err
	}

	workshopData.WorkshopID = workshopID

	return workshopData, nil
}

func visitPage(url string) (*WorkshopData, error) {

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
	} else if len(info.CreatedBy) == 0 {
		return nil, errMissingCreatedBy
	} else if len(info.CreatorName) == 0 {
		return nil, errMissingCreatorName
	} else if len(info.FileSize) == 0 {
		return nil, errMissingFileSize
	} else if len(info.PreviewImage) == 0 {
		return nil, errMissingPreviewImage
	}

	malm.Info("name: %s", info.CreatorName)

	return info, err
}
