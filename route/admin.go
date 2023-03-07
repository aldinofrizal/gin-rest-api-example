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
	restaurantController := controller.RestaurantController{}
	restaurants.Use(middleware.Authentication())
	{
		restaurants.GET("/", restaurantController.Index)
		restaurants.GET("/:id", restaurantController.Detail)
		restaurants.POST("/", restaurantController.Create)
		restaurants.PUT("/:id", restaurantController.Update)
		restaurants.DELETE("/:id", restaurantController.Delete)
	}
}
