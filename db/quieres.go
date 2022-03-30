package db

import (
	"github.com/nexsabre/mikropoker-go/models"
	"gorm.io/gorm"
)

func GetSessions(db *gorm.DB) []models.Session {
	var sessions []models.Session
	db.Find(&sessions)
	return sessions
	// var user models.User
	// db.Create(&models.User{Username: "test", Salle: 3.0, SessionID: int(session.ID)})

	// db.Preload("Session").Last(&user)
	// fmt.Fprintf(os.Stdout, "%+v", user)
}

// func GetSession(db *grom.DB, sessionID: int) {
// 	var session models.Session
// 	db.Create(&models.Session{Name: "Session1", Reveal: true})
// 	db.Last(&session)

// 	var user models.User
// 	db.Create(&models.User{Username: "test", Salle: 3.0, SessionID: int(session.ID)})

// 	db.Preload("Session").Last(&user)
// 	fmt.Fprintf(os.Stdout, "%+v", user)
// }
