package database

import (
	"log"
	"os"
	"github.com/prashantsainii/fiber-REST-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {		// struct to hold out database pointer
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err.Error())
		os.Exit(2)
	}

	log.Println("Successfully connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	// Add migrations- create tables
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}