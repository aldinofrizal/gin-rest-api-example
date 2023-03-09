package models

import (
	"errors"

	"github.com/aldinofrizal/gin-ozamot-api/entity/response"
	"github.com/aldinofrizal/gin-ozamot-api/utilities"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := utilities.HashPassword(u.Password)

	if err != nil {
		err = errors.New("failed to hash password")
	}

	u.Password = hashedPassword
	return
}

func (u *User) GetResponse() response.User {
	return response.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func (u *User) InvalidLogin() []utilities.ApiBindError {
	return []utilities.ApiBindError{
		{
			Field: "email/password",
			Msg:   "Invalid value",
		},
	}
}

func (u *User) IsEmailExist() error {
	userFind := User{}
	result := DB.Where("email = ?", u.Email).First(&userFind)

	if result.RowsAffected > 0 {
		return errors.New("email already exist")
	}

	return nil
}
