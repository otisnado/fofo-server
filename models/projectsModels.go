package models

import "time"

type Project struct {
	ID         uint      `json:"id" gorm:"primary_key; not null"`
	Name       string    `json:"name" gorm:"not null"`
	Created_by string    `json:"created_by" gorm:"not null"`
	Language   string    `json:"language" gorm:"not null"`
	CreatedAt  time.Time `json:"created"`
	UpdatedAt  time.Time `json:"updated"`
}
