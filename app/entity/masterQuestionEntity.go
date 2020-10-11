package entity

import (
	"time"
)

type MasterQuestion struct {
	ID           int       `json:"id" ;gorm:"AUTO_INCREMENT"`
	QuestionText string    `json:"question_text"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	MasterAnswer []MasterAnswer
}

type MasterQuestionCreate struct {
	QuestionText string `json:"question_text"`
}

type MasterQuestionDelete struct {
	ID int `json:"id"`
}
