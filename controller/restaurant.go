package controller

import (
	"net/http"

	"github.com/aldinofrizal/gin-ozamot-api/entity/request"
	"github.com/gin-gonic/gin"
)

type RestaurantController struct {
}

func (r *RestaurantController) Index(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Restaurants Index",
	})
}

func (r *RestaurantController) Detail(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Restaurant Detail",
		"id":      id,
	})
}

func (r *RestaurantController) Create(c *gin.Context) {
	var body request.Restaurant
	c.BindJSON(&body)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Restaurant Create",
		"body":    body,
	})
}

func (r *RestaurantController) Update(c *gin.Context) {
	id := c.Param("id")
	var body request.Restaurant
	c.BindJSON(&body)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Restaurant Update",
		"body":    body,
		"id":      id,
	})
}

func (r *RestaurantController) Delete(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusCreated, gin.H{
		"message": "Restaurant Delete",
		"id":      id,
	})
}
