package entity

import (
	"time"
)

type MasterAnswer struct {
	ID               int       `json:"id" ;gorm:"AUTO_INCREMENT"`
	MasterQuestionId int       `json:"master_question_id"`
	FormAnswer       string    `json:"form_answer"`
	FormScore        string    `json:"form_score"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type MasterAnswerCreate struct {
	MasterQuestionId int    `json:"master_question_id"`
	FormAnswer       string `json:"form_answer"`
	FormScore        string `json:"form_score"`
}

type MasterAnswerDelete struct {
	ID int `json:"id"`
}
