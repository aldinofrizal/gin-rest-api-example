package models

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	ImageUrl    string `gorm:"not null" json:"image_url"`
	AuthorId    uint   `gorm:"not null" json:"author_id"`
	Author      User   `gorm:"foreignKey:AuthorId"`
}
