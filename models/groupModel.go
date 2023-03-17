package models

import "time"

type Group struct {
	ID        uint      `json:"id" gorm:"primary_key; not null; autoIncrement"`
	Name      string    `json:"name" gorm:"not null"  binding:"required"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type GroupUpdate struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type SuccessFindGroups struct {
	Data []Project `json:"data"`
}

type SuccessFindGroup struct {
	Data Project `json:"data"`
}

type SuccessGroupCreation struct {
	Data Project `json:"data"`
}

type SuccessGroupDelete struct {
	Data bool `json:"data"`
}

type SuccessGroupUpdate struct {
	Data Project `json:"data"`
}
