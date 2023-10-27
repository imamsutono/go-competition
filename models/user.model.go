package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Username    string    `gorm:"type:varchar(20);not null"`
	Email       string    `gorm:"uniqueIndex;not null"`
	Password    string    `gorm:"not null"`
	CreatedAt   time.Time
	Competition []*Competition `gorm:"many2many:user_competitions;" json:"competitions"`
}
