package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		// go mailer.RecoveryMail(recovered.(string), c.Request.Host+c.Request.URL.Path)
		if _, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error, we already received your error and will handle it soon!",
			})
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	apiV1 := r.Group("/api/v1")
	SetupAdminRoute(apiV1)

	// customers := r.Group("/pub/v1")

}
