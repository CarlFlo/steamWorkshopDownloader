package server

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/CarlFlo/steamWorkshopDownloader/database"
	steamworkshop "github.com/CarlFlo/steamWorkshopDownloader/steamWorkshop"
)

func getItemPipeline(fi *FileInfo) (*database.WorkshopItem, error) {

	// Look in the cache

	var workshopData *database.WorkshopItem
	workshopData.QueryItemByWorkshopID(fi.WorkshopID)

	/*
		workshopData, err := steamworkshop.ParseWorkshopURL(fi.Url, fi.WorkshopID)
		if err != nil {
			return nil, err
		}
	*/

	item, err := steamworkshop.DownloadWorkshopItem(workshopData.GetCommand())
	if err != nil {
		return nil, err
	}

	// The zipFileName or ID is unique to every workshop file. Can be used for cache
	if err := zipFolder(item.PathToFile, fi.ZipFilePath); err != nil {
		return nil, err
	}

	return workshopData, nil
}

// zipFolder creates a zip archive of the specified folder.
func zipFolder(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == source {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
