package models

import "time"

type Language struct {
	ID         int       `json:"id" gorm:"primary_key; not null"`
	Name       string    `json:"name" gorm:"not null"`
	Created_by uint      `json:"created_by" gorm:"not null"`
	CreatedAt  time.Time `json:"created"`
	UpdatedAt  time.Time `json:"updated"`
}
