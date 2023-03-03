package models

type Project struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Created_by string `json:"created_by"`
	Language   string `json:"language"`
}

type InputProject struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name" binding:"required"`
	Created_by string `json:"created_by" binding:"required"`
	Language   string `json:"language" binding:"required"`
}

type UpdateProjectInput struct {
	Name       string `json:"name"`
	Created_by string `json:"created_by"`
	Language   string `json:"language"`
}
