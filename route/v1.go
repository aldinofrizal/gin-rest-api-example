package route

import (
	"github.com/aldinofrizal/gin-rest-api-example/controller"
	"github.com/aldinofrizal/gin-rest-api-example/middleware"
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

	news := r.Group("/news")
	newsController := controller.NewsController{}
	news.Use(middleware.Authentication())
	{
		news.GET("", newsController.Index)
	}

	tvshows := r.Group("/tvshows")
	tvshowsController := controller.TvshowsController{}
	tvshows.Use(middleware.Authentication())
	{
		tvshows.GET("", tvshowsController.Index)
		tvshows.GET("/:id", tvshowsController.Detail)
	}

	bookmark := r.Group("/bookmarks")
	bookmarkController := controller.BookmarkController{}
	bookmark.Use(middleware.Authentication())
	{
		bookmark.POST("", bookmarkController.Create)
		bookmark.GET("", bookmarkController.Index)
		bookmark.DELETE("/:tmdb_id", bookmarkController.Delete)
	}
}
