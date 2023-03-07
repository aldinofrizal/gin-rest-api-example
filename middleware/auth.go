package middleware

import (
	"github.com/gin-gonic/gin"
)

// type HeaderRequest struct {
// 	access_token string `header:"access_token" binding:"required"`
// }

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from c.Request.Header.Get("access_token")
		// jwt logic
		c.Next()
	}
}
