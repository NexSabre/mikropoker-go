package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Username  string     `json:"username"`
	Salle     float32    `json:"salle"`
	SessionID int        `json:"-"`
}

type Session struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Name      string     `json:"name"`
	Reveal    bool       `json:"reveal"`
	Users     []User     `json:"users"`
}
