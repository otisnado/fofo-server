package models

import "time"

type Language struct {
	ID         int       `json:"id" gorm:"primary_key; not null"`
	Name       string    `json:"name" gorm:"not null"`
	Created_by uint      `json:"created_by" gorm:"not null"`
	CreatedAt  time.Time `json:"created"`
	UpdatedAt  time.Time `json:"updated"`
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
