package controller

import (
	"net/http"

	"github.com/aldinofrizal/gin-ozamot-api/entity/models"
	"github.com/aldinofrizal/gin-ozamot-api/entity/request"
	"github.com/aldinofrizal/gin-ozamot-api/utilities"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct{}

func (u *UserController) Register(c *gin.Context) {
	var registerInput request.Register

	if err := c.ShouldBindJSON(&registerInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
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
		c.AbortWithStatusJSON(http.StatusBadRequest, result.Error.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Regiseter",
		"user":    user.GetResponse(),
	})
}

func (u *UserController) Login(c *gin.Context) {
	var loginUser request.Login

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utilities.ParseError(err),
		})
		return
	}

	userFind := models.User{}
	result := models.DB.Where("email = ?", loginUser.Email).First(&userFind)

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": userFind.InvalidLogin(),
		})
		return
	}

	if !utilities.CheckPasswordHash(loginUser.Password, userFind.Password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": userFind.InvalidLogin(),
		})
		return
	}

	token, _ := utilities.GenerateToken(jwt.MapClaims{"ID": userFind.ID})

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Login",
		"user":    userFind.GetResponse(),
		"token":   token,
	})
}

func (u *UserController) CurrentLoggedUser(c *gin.Context) {
	loggedUser := c.MustGet("user").(*models.User)
	c.JSON(http.StatusOK, gin.H{
		"user": loggedUser.GetResponse(),
	})
}
