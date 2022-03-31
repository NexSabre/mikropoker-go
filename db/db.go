package db

import (
	"os"

	"github.com/nexsabre/mikropoker-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	DATABASE_URL := "DATABASE_URL"
	db, err = gorm.Open(postgres.Open(os.Getenv(DATABASE_URL)), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}
