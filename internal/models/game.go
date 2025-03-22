package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Platform    string `json:"platform"`
}
