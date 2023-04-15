package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aldinofrizal/gin-rest-api-example/entity/models"
	"github.com/aldinofrizal/gin-rest-api-example/entity/request"
	"github.com/aldinofrizal/gin-rest-api-example/utilities"
	"github.com/gin-gonic/gin"
)

type BookmarkController struct{}

func (r *BookmarkController) Create(c *gin.Context) {
	var bookmarkBody request.Bookmark

	if err := c.ShouldBindJSON(&bookmarkBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": utilities.ParseError(err),
		})
		return
	}

	user := c.MustGet("user").(*models.User)
	bookmark := models.Bookmark{
		Name:       bookmarkBody.Name,
		Overview:   bookmarkBody.Overview,
		PosterPath: bookmarkBody.PosterPath,
		TmdbId:     bookmarkBody.TmdbId,
		UserId:     int(user.ID),
	}

	if bookmark.IsExist() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "already exists",
		})
		return
	}

	result := models.DB.Create(&bookmark)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.Error.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Success add %s to your bookmark", bookmark.Name),
	})
}

func (r *BookmarkController) Delete(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	tmdbId := c.Param("tmdb_id")
	tmdbIdInt, _ := strconv.Atoi(tmdbId)

	bookmark := models.Bookmark{
		UserId: int(user.ID),
		TmdbId: tmdbIdInt,
	}

	if !bookmark.IsExist() {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "bookmark not found",
		})
		return
	}

	result := models.DB.Delete(&bookmark)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete bookmark",
	})
}

func (r *BookmarkController) Index(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	bookmarks := []models.Bookmark{}

	result := models.DB.Where("user_id = ?", user.ID).Joins("User").Find(&bookmarks)
	if result.Error != nil {
		panic("db query error")
	}

	c.JSON(http.StatusOK, gin.H{
		"result": bookmarks,
	})
}
