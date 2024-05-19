package server

import (
	"net/http"
	"os"

	"github.com/CarlFlo/malm"
)

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
