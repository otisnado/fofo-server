package models

type TokenRequest struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
