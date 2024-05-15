package main

import (
	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/steamWorkshopDownloader/config"
	"github.com/CarlFlo/steamWorkshopDownloader/server"
)

func init() {

	malm.SetLogVerboseBitmask(0)
	config.Load()
}

func main() {

	server.Launch()
}
