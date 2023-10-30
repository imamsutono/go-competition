package models

import (
	"time"
)

type User struct {
	ID           int64          `gorm:"primaryKey;column:id" json:"registration_id"`
	Username     string         `gorm:"column:username;type:varchar(20);not null" json:"username" validate:"required,alphanum,lowercase,min=3,max=16"`
	Email        string         `gorm:"column:email;type:varchar(255);uniqueIndex;not null" json:"email" validate:"required,email"`
	Password     string         `gorm:"column:password;not null" validate:"required,min=8"`
	CreatedAt    time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"join_date"`
	Competitions []*Competition `gorm:"many2many:user_competitions;" json:"competitions"`
}
