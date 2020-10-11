package entity

import "time"

type MasterForm struct {
	ID        int       `json:"id" ;gorm:"AUTO_INCREMENT"`
	FormName  string    `json:"form_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MasterFormCreate struct {
	FormName string `json:"form_name"`
}

type MasterFormDelete struct {
	ID int `json:"id"`
}
