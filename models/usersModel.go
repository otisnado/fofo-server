package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

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

type SuccessFindUsers struct {
	Data []User `json:"data"`
}

type SuccessFindUser struct {
	Data User `json:"data"`
}

type SuccessUserCreation struct {
	Data User `json:"data"`
}

type SuccessUserDelete struct {
	Data bool `json:"data"`
}

type SuccessUserUpdate struct {
	Data User `json:"data"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
