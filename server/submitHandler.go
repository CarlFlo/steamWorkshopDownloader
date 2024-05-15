package server

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/CarlFlo/malm"
)

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Could not parse form", http.StatusBadRequest)
		malm.Error("could not parse form. %v", err)
		return
	}

	if err := CheckUrlInput(r.FormValue("inputText")); err != nil {
		http.Error(w, "Invalid URL provided", http.StatusBadRequest)
		malm.Error("user provided an invalid URL. %v", err)
		return
	}

	malm.Debug("Processing submit request")

	fi := &FileInfo{}
	fi.New(r.FormValue("inputText"))

	// Check if the file is already downloaded and zipped
	if !lookForFileOnDisk(fi) {
		// No zip file with that ID in the cache
		_, err := prepareWorkshopItem(fi)
		if err != nil {
			http.Error(w, "Something went wrong.", http.StatusInternalServerError)
			malm.Error("could not fetch information about the workshop item. %v", err)
			return
		}
	}

	file, err := os.Open(fi.ZipFilePath)
	if err != nil {
		http.Error(w, "File not found.", http.StatusNotFound)
		malm.Error("could not open the file. %v", err)
		return
	}
	defer file.Close()

	// Set the headers
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fi.ZipFileName))
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize(fi.ZipFilePath)))

	// Stream the file content to the response
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Error serving file.", http.StatusInternalServerError)
	}

}

func fileSize(filepath string) int64 {
	info, err := os.Stat(filepath)
	if err != nil {
		return -1
	}
	return info.Size()
}
