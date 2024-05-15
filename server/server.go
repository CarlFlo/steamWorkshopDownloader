package server

import (
	"fmt"
	"net/http"
	"os"

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

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the HTML file
	html, err := os.ReadFile("static/index.html")
	if err != nil {
		http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
		malm.Error("%v", err)
		return
	}

	// Set the Content-Type header and write the HTML content
	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}
