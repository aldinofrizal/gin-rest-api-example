package controller

import (
	"net/http"

	"github.com/aldinofrizal/gin-ozamot-api/entity/request"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u *UserController) Register(c *gin.Context) {
	var body request.Register

	c.BindJSON(&body)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Regiseter",
		"body":    body,
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
