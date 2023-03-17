package models

import "time"

type Language struct {
	ID         int       `json:"id" gorm:"primary_key; not null"`
	Name       string    `json:"name" gorm:"not null"  binding:"required"`
	Created_by uint      `json:"created_by" gorm:"not null"  binding:"required"`
	CreatedAt  time.Time `json:"created"`
	UpdatedAt  time.Time `json:"updated"`
}

type LanguageUpdate struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Created_by uint   `json:"created_by,omitempty"`
}

type SuccessFindLanguages struct {
	Data []Project `json:"data"`
}

type SuccessFindLanguage struct {
	Data Project `json:"data"`
}

type SuccessLanguageCreation struct {
	Data Project `json:"data"`
}

type SuccessLanguageDelete struct {
	Data bool `json:"data"`
}

type SuccessLanguageUpdate struct {
	Data Project `json:"data"`
}
