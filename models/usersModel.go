package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key; not null"`
	Name      string    `json:"name" binding:"required"`
	Lastname  string    `json:"lastname" binding:"required"`
	Username  string    `json:"username" gorm:"not null; unique" binding:"required"`
	Mail      string    `json:"mail" gorm:"not null; unique" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	Group     int       `json:"group" binding:"required"`
	State     bool      `json:"state" binding:"required"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type UserUpdate struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Username string `json:"username,omitempty"`
	Mail     string `json:"mail,omitempty"`
	Password string `json:"password,omitempty"`
	Role     uint   `json:"roles,omitempty"`
	Group    int    `json:"group,omitempty"`
	State    bool   `json:"state,omitempty"`
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
