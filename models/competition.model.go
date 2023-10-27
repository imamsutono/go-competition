package models

import "time"

type Competition struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"varchar(255);not null" json:"name"`
	StartDate string    `gorm:"varchar(20);not null" json:"start_date"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	User      []*User   `gorm:"many2many:user_competitions;" json:"participants"`
}

type CompetitionPostRequest struct {
	Name      string `json:"name,omitempty"`
	StartDate string `json:"startDate,omitempty"`
}
