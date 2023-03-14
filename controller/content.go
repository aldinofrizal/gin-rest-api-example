package controller

import (
	"fmt"
	"net/http"

	"github.com/aldinofrizal/gin-ozamot-api/entity/models"
	"github.com/aldinofrizal/gin-ozamot-api/entity/request"
	"github.com/aldinofrizal/gin-ozamot-api/utilities"
	"github.com/gin-gonic/gin"
)

type ContentController struct {
}

func (r *ContentController) Index(c *gin.Context) {
	// loggedUser := c.MustGet("user").(*models.User)
	contents := []models.Content{}
	result := models.DB.Joins("Author").Find(&contents)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Content Index",
		"contents": contents,
	})
}

func (r *ContentController) Detail(c *gin.Context) {
	id := c.Param("id")
	content := models.Content{}
	result := models.DB.Joins("Author").First(&content, id)

	if result.Error != nil || result.RowsAffected == 0 {
		fmt.Println(result.Error.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Content Detail",
		"content": content,
	})
}

func (r *ContentController) Create(c *gin.Context) {
	loggedUser := c.MustGet("user").(*models.User)
	var newContent request.Content

	if err := c.ShouldBindJSON(&newContent); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": utilities.ParseError(err),
		})
		return
	}

	content := models.Content{
		Name:        newContent.Name,
		Description: newContent.Description,
		ImageUrl:    newContent.ImageUrl,
		AuthorId:    loggedUser.ID,
	}

	result := models.DB.Create(&content)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.Error.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Content Created",
		"content": content,
	})
}

func (r *ContentController) Update(c *gin.Context) {
	id := c.Param("id")
	var updateData request.Content
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": utilities.ParseError(err),
		})
		return
	}

	result := models.DB.Model(&models.Content{}).Where("id = ?", id).Updates(&models.Content{
		Name:        updateData.Name,
		Description: updateData.Description,
		ImageUrl:    updateData.ImageUrl,
	})

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "resource not found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Content Update",
		"body":    updateData,
		"id":      id,
	})
}

func (r *ContentController) Delete(c *gin.Context) {
	content := c.MustGet("contentToDelete").(*models.Content)
	result := models.DB.Delete(&content)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Content Deleted",
		"id":      content.ID,
	})
}
