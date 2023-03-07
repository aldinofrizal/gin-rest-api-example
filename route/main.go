package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {

	apiV1 := r.Group("/api/v1")
	SetupAdminRoute(apiV1)

	// customers := r.Group("/pub/v1")
}
