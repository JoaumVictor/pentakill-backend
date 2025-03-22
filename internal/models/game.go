package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Name        string          `gorm:"not null" json:"name"`
	Description string          `json:"description"`
	Assessment  int             `gorm:"not null; check:assessment >= 0 AND assessment <= 10" json:"assessment"`
	Genre       string          `json:"genre"`
	Platform    string          `gorm:"not null; check:platform IN ('Steam', 'Xbox', 'Playstation')" json:"platform"`
	Date        time.Time       `gorm:"not null" json:"date"`
	Finished    bool            `json:"finished"`
	Price       float64         `json:"price"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"-"`
}