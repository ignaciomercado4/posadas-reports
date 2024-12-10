package database

import (
	"better-posadas/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("succesful database connection")
	return db
}

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.Report{})
	fmt.Println("models migrated")
}
