package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
	// CreatedAt time.Time      `json:"created_at"`
	// UpdatedAt time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
