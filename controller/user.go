package controller

import (
	"net/http"

	"github.com/aldinofrizal/gin-ozamot-api/entity/models"
	"github.com/aldinofrizal/gin-ozamot-api/entity/request"
	"github.com/aldinofrizal/gin-ozamot-api/utilities"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) Register(c *gin.Context) {
	var registerInput request.Register

	if err := c.ShouldBindJSON(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utilities.ParseError(err),
		})
		return
	}

	user := models.User{
		Name:     registerInput.Name,
		Email:    registerInput.Email,
		Password: registerInput.Password,
	}

	if err := user.IsEmailExist(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": []utilities.ApiBindError{
				{
					Field: "Email",
					Msg:   err.Error(),
				},
			},
		})
		return
	}

	result := models.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Regiseter",
		"user":    user.GetResponse(),
	})
}

func (u *UserController) Login(c *gin.Context) {
	var body request.Login

	c.BindJSON(&body)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Login",
		"body":    body,
	})
}
