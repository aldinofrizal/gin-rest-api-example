package models

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	ImageUrl    string `gorm:"not null"`
	AuthorId    uint   `gorm:"not null"`
}
