package db

import (
	"github.com/nexsabre/mikropoker-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	// DATABASE_URL := "DATABASE_URL"
	// db, err = gorm.Open(postgres.Open(os.Getenv(DATABASE_URL)), &gorm.Config{})
	db, err = gorm.Open(sqlite.Open("sql_app.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect db")
	}

	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}
