package controller

import (
	"net/http"
	"strconv"

	"github.com/aldinofrizal/gin-rest-api-example/services/tmdb"
	"github.com/gin-gonic/gin"
)

type TvshowsController struct{}

func (r *TvshowsController) Index(c *gin.Context) {
	queryPage := c.Query("page")
	if queryPage == "" {
		queryPage = "1"
	}
	tmdb := tmdb.ImplTmdbClient()
	resp, err := tmdb.GetMovies(queryPage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"messages": "Something went wrong, try again later",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": resp.Results,
		"page":    queryPage,
	})
}

func (r *TvshowsController) Detail(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	tmdb := tmdb.ImplTmdbClient()
	resp, err := tmdb.Detail(intId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"messages": "Something went wrong, try again later",
			"error":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": resp,
	})
}
