package models

import "time"

type Project struct {
	ID         uint      `json:"id" gorm:"primary_key; not null"`
	Name       string    `json:"name" gorm:"not null"  binding:"required"`
	Created_by string    `json:"created_by" gorm:"not null"  binding:"required"`
	Language   string    `json:"language" gorm:"not null"  binding:"required"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}

type ProjectUpdate struct {
	ID         uint   `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Created_by string `json:"created_by,omitempty"`
	Language   string `json:"language,omitempty"`
}

type SuccessFindProjects struct {
	Data []Project `json:"data"`
}

type SuccessFindProject struct {
	Data Project `json:"data"`
}

type SuccessProjectCreation struct {
	Data Project `json:"data"`
}

type SuccessProjectDelete struct {
	Data bool `json:"data"`
}

type SuccessProjectUpdate struct {
	Data Project `json:"data"`
}
