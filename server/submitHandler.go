package server

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/steamWorkshopDownloader/database"
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

	var workshopData *database.WorkshopItem

	// Check if the file is already downloaded and zipped
	if checkCache(fi) {
		// Load cached workshopData data
		workshopData = LoadWorkshopData(fi.WorkshopID)
	} else {
		// No zip file with that ID in the cache
		var err error
		// TODO: Remake function getting info about the item and downloading the item should be seperated. Because we want to display info about the workshop url as fast as possible to the user
		workshopData, err = getItemPipeline(fi)
		if err != nil {
			http.Error(w, "Soemthing went wrong.", http.StatusInternalServerError)
			malm.Error("could not fetch information about the workshop item. %v", err)
			return
		}
	}

	// So VSC stops complaining about unused variables...
	workshopData.GetCommand()

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
