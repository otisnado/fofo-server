package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key; not null"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Username  string    `json:"username" gorm:"not null; unique"`
	Mail      string    `json:"mail" gorm:"not null; unique"`
	Password  string    `json:"password"`
	Group     int       `json:"group"`
	State     bool      `json:"state"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
