package models

import (
	"time"
)

type Competition struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	Name      string    `gorm:"column:name;varchar(255);not null" json:"name" validate:"required,min=3,max=255"`
	StartDate string    `gorm:"column:start_date;varchar(20);not null" json:"startDate" validate:"required,datetime"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	Users     []*User   `gorm:"many2many:user_competitions;" json:"participants"`
}

type CompetitionPostRequest struct {
	Name      string `json:"name,omitempty"`
	StartDate string `json:"startDate,omitempty"`
}

type CompetitionsPost201Response struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	StartDate string `json:"startDate,omitempty"`
}

type CompetitionsIdGet200Response struct {
	Id           *int64                                          `json:"id,omitempty"`
	Name         *string                                         `json:"name,omitempty"`
	Participants *[]CompetitionsIdGet200ResponseParticipantsItem `json:"participants,omitempty"`
	StartDate    *string                                         `json:"startDate,omitempty"`
}

type CompetitionsIdGet200ResponseParticipantsItem struct {
	RegistrationId int64  `json:"registration_id"`
	Username       string `json:"username"`
}
