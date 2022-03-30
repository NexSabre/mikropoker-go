package db

import (
	"fmt"

	"github.com/nexsabre/mikropoker-go/models"
	"github.com/nexsabre/mikropoker-go/schema"
	"gorm.io/gorm"
)

func GetSessions(db *gorm.DB) []models.Session {
	var sessions []models.Session
	db.Find(&sessions)
	return sessions

}

func GetSession(db *gorm.DB, sessionID int) models.Session {
	var session models.Session

	db.Preload("Users").First(&session, sessionID)
	return session
}

func CreateSession(db *gorm.DB, session_body schema.SessionCreate) models.Session {
	newSession := models.Session{
		Name:   session_body.Name,
		Reveal: false,
	}

	db.Create(&newSession)
	db.Last(&newSession)
	return newSession
}

func RevealSession(db *gorm.DB, sessionID int, reveal bool) models.Session {
	sessionStatus := models.Session{
		Reveal: reveal,
	}
	db.Model(&sessionStatus).Where("id = ?", sessionID).Updates(sessionStatus)
	db.Find(&sessionStatus, sessionID)
	return sessionStatus
}

func UserPoints(db *gorm.DB, session_id int, user_body schema.UserPoints) models.User {
	userPoints := models.User{
		Username:  user_body.Username,
		Salle:     user_body.Salle,
		SessionID: session_id,
	}

	fmt.Printf("%+v", user_body)
	if db.Model(&userPoints).
		Where("username = ?", user_body.Username).
		Updates(&userPoints).RowsAffected == 0 {
		fmt.Println("0asdas")
		db.Create(&userPoints)
	}
	fmt.Println("asdf")
	db.Last(&userPoints)
	return userPoints
}
