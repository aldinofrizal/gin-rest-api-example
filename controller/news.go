package controller

import (
	"context"
	"net/http"

	"github.com/aldinofrizal/gin-rest-api-example/services"
	pb "github.com/aldinofrizal/gin-rest-api-example/services/news"
	"github.com/gin-gonic/gin"
)

type NewsController struct{}

func (r *NewsController) Index(c *gin.Context) {
	n, err := services.NewsService()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": "Something went wrong, try again later",
		})
		return
	}

	news, err := n.GetNews(context.Background(), &pb.EmptyParams{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": "Something went wrong, try again later",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"news":    news.Items,
	})
}
