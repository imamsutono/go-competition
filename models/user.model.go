package models

import (
	"time"
)

type User struct {
	ID           int64          `gorm:"primaryKey" json:"registration_id"`
	Username     string         `gorm:"type:varchar(20);not null" json:"username" validate:"required,alphanum,lowercase,min=3,max=16"`
	Email        string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email" validate:"required,email"`
	Password     string         `gorm:"not null" validate:"required,min=8"`
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"join_date"`
	Competitions []*Competition `gorm:"many2many:user_competitions;" json:"competitions"`
}
