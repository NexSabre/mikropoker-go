package db

import (
	"github.com/nexsabre/mikropoker-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB_NAME string = "sql_app.db"
var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}
