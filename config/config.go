package config

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/CarlFlo/malm"
)

// CONFIG holds all the config data
var CONFIG *configStruct

type configStruct struct {
	Port           int      `json:"port"`
	PathToSteamCMD string   `json:"pathToSteamCMD"`
	PathToCache    string   `json:"pathToCache"`
	Database       database `json:"database"`
}

type database struct {
	FileName string `json:"fileName"`
}

// ReloadConfig is a wrapper function for reloading the config. For clarity
func ReloadConfig() error {
	return readConfig()
}

// readConfig will read the config file
func readConfig() error {

	file, err := os.Open("./config.json")
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	file.Close()

	if err = json.Unmarshal(buf.Bytes(), &CONFIG); err != nil {
		return err
	}

	return nil
}

// createConfig creates the default config file
func createConfig() error {

	// Default config settings
	configStruct := configStruct{
		Port:           8080,
		PathToSteamCMD: "",
		PathToCache:    "",
		Database: database{
			FileName: "database.db",
		},
	}

	jsonData, _ := json.MarshalIndent(configStruct, "", "   ")
	err := os.WriteFile("config.json", jsonData, 0644)

	return err
}

// loadConfiguration loads the configuration file
func loadConfiguration() error {

	if err := readConfig(); err != nil {
		if err = createConfig(); err != nil {
			return err
		}
		if err = readConfig(); err != nil {
			return err
		}
	}
	return nil
}

func Load() {
	if err := loadConfiguration(); err != nil {
		malm.Fatal("Error loading configuration: %v", err)
		return
	}

	requiredVariableCheck()

	malm.Info("Configuration loaded")
}

// Some variables are required for the bot to work
func requiredVariableCheck() {

	// This function checks if some important variables are set in the config file
	problem := false

	// Must be over 0 and not over 65535
	if CONFIG.Port <= 0 || CONFIG.Port > 65535 {
		malm.Error("No valid port provided")
		problem = true
	}

	if len(CONFIG.PathToSteamCMD) == 0 {
		malm.Error("No path to SteamCMD.exe provided")
		problem = true
	}

	if len(CONFIG.PathToCache) == 0 {
		malm.Error("No path to the filecache provided")
		problem = true
	}

	// Enter more checks here as needed

	if problem {
		malm.Fatal("There are at least one variable missing in the configuration file. Please fix the above errors!")
	}
}
