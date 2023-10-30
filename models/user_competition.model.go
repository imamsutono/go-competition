package models

type UserCompetition struct {
	RegistrationID int64 `gorm:"primaryKey" json:"registration_id"`
	CompetitionID  int64 `gorm:"foreignKey"`
	UserID         int64 `gorm:"foreignKey"`
}

type UserCompetitionJSONRequestBody struct {
	CompetitionID string `json:"competition,omitempty"`
	UserID        int    `json:"userid,omitempty"`
}
