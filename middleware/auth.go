package middleware

import (
	"net/http"

	"github.com/aldinofrizal/gin-rest-api-example/entity/models"
	"github.com/aldinofrizal/gin-rest-api-example/utilities"
	"github.com/gin-gonic/gin"
)

// type HeaderRequest struct {
// 	access_token string `header:"access_token" binding:"required"`
// }

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
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
