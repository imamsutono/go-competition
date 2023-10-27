package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key" json:"registration_id"`
	Username     string    `gorm:"type:varchar(20);not null" json:"username"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	Password     string    `gorm:"not null"`
	CreatedAt    time.Time
	Competitions []*Competition `gorm:"many2many:user_competitions;" json:"competitions"`
}
