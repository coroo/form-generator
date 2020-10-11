package entity

import (
	"time"
)

type FormGenerator struct {
	ID               int            `json:"id" ;gorm:"AUTO_INCREMENT"`
	MasterQuestionId int            `json:"master_question_id"`
	MasterFormId     int            `json:"master_form_id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	MasterQuestion   MasterQuestion `gorm:"foreignkey:master_question_id;AssociationForeignKey:ID"`
	MasterForm       MasterForm     `gorm:"foreignkey:master_form_id;AssociationForeignKey:ID"`
}

type FormGeneratorCreate struct {
	MasterQuestionId int `json:"master_question_id"`
	MasterFormId     int `json:"master_form_id"`
}

type FormGeneratorDelete struct {
	ID int `json:"id"`
}
