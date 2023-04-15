package models

import (
	"github.com/aldinofrizal/gin-rest-api-example/entity/response"
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name"`
	Overview   string `gorm:"not null" json:"overview"`
	TmdbId     int    `gorm:"not null" json:"tmdb_id"`
	PosterPath string `gorm:"not null" json:"poster_path"`
	UserId     int    `gorm:"not null" json:"user_id"`
	User       response.User
}

func (b *Bookmark) IsExist() (exist bool) {
	result := DB.Where("user_id = ? AND tmdb_id = ?", b.UserId, b.TmdbId).First(&b)
	return result.RowsAffected > 0
}
