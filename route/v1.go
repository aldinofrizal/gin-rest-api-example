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
		users.GET("/me", middleware.Authentication(), userController.CurrentLoggedUser)
		users.GET("/verify", userController.Verify)
	}

	contents := r.Group("/contents")
	contentController := controller.ContentController{}
	contents.Use(middleware.Authentication())
	{
		contents.GET("", contentController.Index)
		contents.GET("/:id", contentController.Detail)
		contents.POST("", contentController.Create)
		contents.PUT("/:id", contentController.Update)
		contents.DELETE("/:id",
			middleware.ContentDeleteAuthorization(),
			contentController.Delete,
		)
	}
}
