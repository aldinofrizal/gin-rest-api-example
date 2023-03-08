package main

import (
	"github.com/aldinofrizal/gin-ozamot-api/entity/models"
	"github.com/aldinofrizal/gin-ozamot-api/route"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	models.DBConnect()
	route.SetupRoute(r)

	r.Run()
}
