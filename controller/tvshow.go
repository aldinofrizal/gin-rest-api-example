package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aldinofrizal/gin-rest-api-example/services/tmdb"
	"github.com/aldinofrizal/gin-rest-api-example/utilities"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type TvshowsController struct{}

func (r *TvshowsController) Index(c *gin.Context) {
	queryPage := c.Query("page")
	if queryPage == "" {
		queryPage = "1"
	}

	keys := fmt.Sprintf("imdb_page_%s", queryPage)
	cached, err := utilities.RDB.Get(keys)

	if err == nil {
		movieList := tmdb.MovieList{}
		_ = json.Unmarshal([]byte(cached), &movieList.Results)

		c.JSON(http.StatusOK, gin.H{
			"results": movieList.Results,
			"page":    queryPage,
		})
		return
	}

	tmdb := tmdb.ImplTmdbClient()
	resp, err := tmdb.GetMovies(queryPage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"messages": "Something went wrong, try again later",
		})
		return
	}

	stringifyData, _ := json.Marshal(resp.Results)
	utilities.RDB.Set(keys, string(stringifyData), 5*time.Minute)

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
