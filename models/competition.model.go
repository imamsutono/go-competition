package models

import "time"

type Competition struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"varchar(255);not null" json:"name" validate:"required,min=3,max=255"`
	StartDate time.Time `gorm:"varchar(20);not null" json:"startDate" validate:"required,datetime"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	Users     []*User   `gorm:"many2many:user_competitions;" json:"participants"`
}

type CompetitionPostRequest struct {
	Name      string `json:"name,omitempty"`
	StartDate string `json:"startDate,omitempty"`
}
