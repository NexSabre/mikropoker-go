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
	sessionStatus := &models.Session{}
	db.Model(&models.Session{}).Where("id = ?", sessionID).Update("reveal", reveal)
	db.Find(&sessionStatus, sessionID)
	return *sessionStatus
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
		db.Create(&userPoints)
	}
	db.Last(&userPoints)
	return userPoints
}

func UserPointsForUser(db *gorm.DB, sessionID int, username string, salle float32) {
	userPoints := models.User{
		Username:  username,
		Salle:     salle,
		SessionID: sessionID,
	}

	if db.Model(&userPoints).
		Where("username = ?", username).
		Updates(&userPoints).RowsAffected == 0 {
		db.Create(&userPoints)
	}
}

func RestartSession(db *gorm.DB, sessionID int) {
	db.Where("session_id = ?", sessionID).Delete(&models.User{})
}
