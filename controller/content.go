package controller

import (
	"net/http"

	"github.com/aldinofrizal/gin-ozamot-api/entity/request"
	"github.com/aldinofrizal/gin-ozamot-api/utilities"
	"github.com/gin-gonic/gin"
)

type ContentController struct {
}

func (r *ContentController) Index(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Content Index",
	})
}

func (r *ContentController) Detail(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Content Detail",
		"id":      id,
	})
}

func (r *ContentController) Create(c *gin.Context) {
	var newContent request.Content

	if err := c.ShouldBindJSON(&newContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utilities.ParseError(err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Content Create",
		"body":    newContent,
	})
}

func (r *ContentController) Update(c *gin.Context) {
	id := c.Param("id")
	var body request.Content
	c.BindJSON(&body)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Content Update",
		"body":    body,
		"id":      id,
	})
}

func (r *ContentController) Delete(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusCreated, gin.H{
		"message": "Content Delete",
		"id":      id,
	})
}
