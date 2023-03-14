package middleware

import (
	"fmt"
	"net/http"

	"github.com/aldinofrizal/gin-ozamot-api/entity/models"
	"github.com/aldinofrizal/gin-ozamot-api/utilities"
	"github.com/gin-gonic/gin"
)

// type HeaderRequest struct {
// 	access_token string `header:"access_token" binding:"required"`
// }

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from c.Request.Header.Get("access_token")
		token := c.Request.Header.Get("access_token")

		claims, err := utilities.DecodeToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Please provide valid access token in your headers",
			})
			return
		}

		user := models.User{}
		result := models.DB.First(&user, claims["ID"])

		if result.Error != nil || result.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Please provide valid access token in your headers",
			})
			return
		}

		c.Set("user", &user)
		c.Next()
	}
}

func ContentDeleteAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedUser := c.MustGet("user").(*models.User)
		contentToDelete := models.Content{}
		result := models.DB.First(&contentToDelete, c.Param("id"))
		if result.Error != nil || result.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data not found",
			})
			return
		}

		fmt.Println("----", contentToDelete.AuthorId, "00000", loggedUser.ID)
		if contentToDelete.AuthorId != loggedUser.ID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Please provide valid access token in your headers",
			})
			return
		}

		c.Set("contentToDelete", &contentToDelete)
		c.Next()
	}
}
