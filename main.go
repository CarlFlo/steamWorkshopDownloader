package main

import (
	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/steamWorkshopDownloader/config"
	"github.com/CarlFlo/steamWorkshopDownloader/database"
	"github.com/CarlFlo/steamWorkshopDownloader/server"
	"github.com/CarlFlo/steamWorkshopDownloader/utils"
)

const CurrentVersion = "2024-05-15"

func init() {

	malm.SetLogVerboseBitmask(0)

	utils.Clear()
	config.Load()
	database.Connect()

	go utils.CheckVersion(CurrentVersion)
}

func main() {

	server.Launch()
}
