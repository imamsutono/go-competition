package models

type UserCompetition struct {
	RegistrationID int64 `gorm:"primaryKey;column:registration_id" json:"registration_id"`
	CompetitionID  int64 `gorm:"foreignKey;column:competition_id"`
	UserID         int64 `gorm:"foreignKey;column:user_id"`
}

type UserCompetitionJSONRequestBody struct {
	CompetitionID string `json:"competition,omitempty"`
	UserID        int    `json:"userid,omitempty"`
}
