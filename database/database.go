package database

import (
	"fmt"
	"os"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/steamWorkshopDownloader/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

const resetDatabaseOnStart = false

func Connect() {
	if err := connectToDB(); err != nil {
		malm.Fatal("Database initialization error: %s", err)
		return
	}
	malm.Info("Connected to database")
}

func connectToDB() error {

	if _, err := os.Stat(config.CONFIG.Database.FileName); os.IsNotExist(err) || resetDatabaseOnStart {
		malm.Info("Could not find the SQLite database. Recreating...")
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(config.CONFIG.Database.FileName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	var modelList = []interface{}{
		&WorkshopItem{},
	}

	if resetDatabaseOnStart {

		malm.Info("Resetting database...")

		type tmp interface {
			TableName() string
		}

		for _, e := range modelList {
			table := e.(tmp).TableName()
			DB.Exec(fmt.Sprintf("DROP TABLE %s", table))
		}
	}

	// Remeber to add new tables to the tableList and not just here!
	return DB.AutoMigrate(modelList...)
}
