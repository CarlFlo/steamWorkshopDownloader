package server

import (
	"fmt"
	"net/http"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/steamWorkshopDownloader/config"
)

func Launch() {

	// Handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/getItemInfo", getItemInfoHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Launch
	malm.Debug("Listening and available on 'http://127.0.0.1:%d/'", config.CONFIG.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.CONFIG.Port), nil); err != nil {
		malm.Fatal("Unable to start server: %v", err)
	}
}
