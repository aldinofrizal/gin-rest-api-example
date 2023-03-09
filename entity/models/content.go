package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Address      string `gorm:"not null"`
	ImageUrl     string `gorm:"not null;column: image_url"`
	InstagramUrl string `gorm:"not null;column: instagram_url"`
	AuthorId     uint   `gorm:"not null;column: author_id"`
}
