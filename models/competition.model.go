package models

import (
	"time"

	"github.com/google/uuid"
)

type Competition struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string    `gorm:"varchar(255);not null"`
	StartDate time.Time
	CreatedAt time.Time
	User      []*User `gorm:"many2many:user_competitions;" json:"participants"`
}
