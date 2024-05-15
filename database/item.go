package database

import (
	"fmt"

	"gorm.io/gorm"
)

type WorkshopItem struct {
	gorm.Model
	AppID      string `json:"app_id" gorm:"uniqueIndex"`
	WorkshopID string `json:"workshop_id"`

	ItemName     string `json:"item_name"`
	LastUpdated  string `json:"last_updated"`
	CreatorLink  string `json:"creator_link"`
	CreatorName  string `json:"creator_name"`
	FileSize     string `json:"file_size"`
	PreviewImage string `json:"preview_image"`
}

func (item *WorkshopItem) GetCommand() string {

	return fmt.Sprintf("workshop_download_item %s %s", item.AppID, item.WorkshopID)
}

func (WorkshopItem) TableName() string {
	return "items"
}

func (i *WorkshopItem) AfterCreate(tx *gorm.DB) error {
	// Log in debug DB maybe
	return nil
}

// Saves the data to the database
// If value doesn't contain a matching primary key, value is inserted.
func (i *WorkshopItem) Save() error {
	return DB.Save(&i).Error
}

func (i *WorkshopItem) DoesItemExist(WorkshopID string) bool {

	var count int
	DB.Raw("SELECT COUNT(*) FROM items WHERE workshop_id = ?", WorkshopID).Scan(&count)
	return count == 1
}
