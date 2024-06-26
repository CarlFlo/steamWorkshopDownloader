package server

import (
	"encoding/json"
	"net/http"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/steamWorkshopDownloader/database"
	"github.com/CarlFlo/steamWorkshopDownloader/server/middleware"
	steamworkshop "github.com/CarlFlo/steamWorkshopDownloader/steamWorkshop"
)

func getItemInfoHandler(w http.ResponseWriter, r *http.Request) {

	// Parse the input value from the request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	inputURL := r.FormValue("inputText")
	if inputURL == "" {
		http.Error(w, "Could not parse form", http.StatusBadRequest)
		malm.Error("could not parse form. %v", err)
		return
	}

	if err := middleware.CheckUrlInput(inputURL); err != nil {
		http.Error(w, "Invalid input provided", http.StatusBadRequest)
		malm.Error("user provided an invalid URL. %v", err)
		return
	}

	malm.Debug("Processing getItemInfo request")

	fi := &FileInfo{}
	fi.New(r.FormValue("inputText"))

	// Check cache
	var workshopData database.WorkshopItem
	if workshopData.DoesItemExist(fi.WorkshopID) {

		workshopData.QueryItemByWorkshopID(fi.WorkshopID)
	} else {
		workshopDataPtr, err := steamworkshop.ParseWorkshopURL(fi.Url, fi.WorkshopID)
		if err != nil {
			malm.Error("%v", err)
			return
		}

		// Save object to database
		workshopData = *workshopDataPtr
		workshopData.Save()
	}

	// Marshal the struct into JSON
	jsonData, err := json.Marshal(workshopData)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the JSON to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
