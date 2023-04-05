package controller

import (
	"net/http"

	"github.com/aldinofrizal/gin-rest-api-example/services/tmdb"
	"github.com/gin-gonic/gin"
)

type TvshowsController struct{}

func (r *TvshowsController) Index(c *gin.Context) {
	tmdb := tmdb.ImplTmdbClient()
	resp, err := tmdb.GetMovies()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"messages": "Something went wrong, try again later",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": resp.Results,
	})
}
