package database

import (
	"fmt"
	"os"

	"github.com/lgmztech/profanity-filter/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connection Opened to Database")
}

func SyncDB() {
	DB.AutoMigrate(
		&models.RateLimitedModel{},
	)
}
