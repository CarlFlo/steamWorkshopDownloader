package steamworkshop

import (
	"strings"

	"github.com/CarlFlo/steamWorkshopDownloader/database"
	"github.com/gocolly/colly"
)

func getWorkshopPageInfo(url string) (*database.WorkshopItem, error) {
	c := colly.NewCollector()

	itemInfo := &database.WorkshopItem{}

	// Selector for AppID
	c.OnHTML("#ig_bottom > div.breadcrumbs > a:nth-child(1)", func(e *colly.HTMLElement) {
		linkToGame := e.Attr("href")
		split := strings.Split(linkToGame, "/")
		itemInfo.AppID = split[len(split)-1]
	})

	// Selector for ItemName
	c.OnHTML("#mainContents > div.workshopItemDetailsHeader > div.workshopItemTitle", func(e *colly.HTMLElement) {
		itemInfo.ItemName = e.Text
	})

	// Selector for LastUpdated
	c.OnHTML("div.detailsStatsContainerRight > div:nth-child(3)", func(e *colly.HTMLElement) {
		itemInfo.LastUpdated = e.Text
	})

	// Selector for CreatedBy
	c.OnHTML("div.rightDetailsBlock > div > div > a", func(e *colly.HTMLElement) {
		itemInfo.CreatorLink = e.Attr("href")
	})

	//
	// Selector for CreatorName
	c.OnHTML("#ig_bottom > div.breadcrumbs > a:nth-child(5)", func(e *colly.HTMLElement) {
		itemInfo.CreatorName = strings.TrimSuffix(e.Text, "'s Workshop")
	})

	// Selector for FileSize
	c.OnHTML("div.detailsStatsContainerRight > div:nth-child(1)", func(e *colly.HTMLElement) {
		itemInfo.FileSize = e.Text
	})

	// Selector for PreviewImage
	c.OnHTML("#previewImageMain", func(e *colly.HTMLElement) {
		previewUrl := e.Attr("src")
		itemInfo.PreviewImage = strings.Split(previewUrl, "?")[0]
	})

	// Backup for PreviewImage above
	c.OnHTML("#previewImage", func(e *colly.HTMLElement) {
		if len(itemInfo.PreviewImage) == 0 {
			previewUrl := e.Attr("src")
			itemInfo.PreviewImage = strings.Split(previewUrl, "?")[0]
		}
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return itemInfo, nil
}
