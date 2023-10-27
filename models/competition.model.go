package models

import (
	"time"

	"github.com/google/uuid"
)

type Competition struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name      string    `gorm:"varchar(255);not null"`
	StartDate string    `gorm:"varchar(20);not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	User      []*User   `gorm:"many2many:user_competitions;" json:"participants"`
}

type CompetitionPostRequest struct {
	Name      string `json:"name,omitempty"`
	StartDate string `json:"startDate,omitempty"`
}
