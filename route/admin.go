package route

import (
	"github.com/aldinofrizal/gin-ozamot-api/controller"
	"github.com/aldinofrizal/gin-ozamot-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupAdminRoute(r *gin.RouterGroup) {
	users := r.Group("/users")
	userController := controller.UserController{}
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
	}

	restaurants := r.Group("/restaurants")
	contentController := controller.ContentController{}
	restaurants.Use(middleware.Authentication())
	{
		restaurants.GET("/", contentController.Index)
		restaurants.GET("/:id", contentController.Detail)
		restaurants.POST("/", contentController.Create)
		restaurants.PUT("/:id", contentController.Update)
		restaurants.DELETE("/:id", contentController.Delete)
	}
}
