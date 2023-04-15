package models

import (
	"errors"

	"github.com/aldinofrizal/gin-rest-api-example/entity/response"
	"github.com/aldinofrizal/gin-rest-api-example/services/mailer"
	"github.com/aldinofrizal/gin-rest-api-example/utilities"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"not null"`
	Email            string `gorm:"uniqueIndex;not null"`
	Password         string `gorm:"not null"`
	VerificationCode string
	IsActive         bool
	Bookmarks        []Bookmark
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := utilities.HashPassword(u.Password)

	if err != nil {
		err = errors.New("failed to hash password")
	}

	u.Password = hashedPassword
	u.IsActive = false
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	generatedToken, _ := utilities.GenerateToken(jwt.MapClaims{"ID": u.ID, "IsVerification": true})
	u.VerificationCode = generatedToken
	go func() {
		DB.Save(&u)
		mailer.RegisterMail(u.Email, u.VerificationCode)
	}()
	return
}

func (u *User) GetResponse() response.User {
	return response.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		IsActive: u.IsActive,
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
