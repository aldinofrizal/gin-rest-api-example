package main

import (
	"github.com/aldinofrizal/gin-rest-api-example/entity/models"
	"github.com/aldinofrizal/gin-rest-api-example/route"
	"github.com/aldinofrizal/gin-rest-api-example/services/mailer"
	"github.com/aldinofrizal/gin-rest-api-example/utilities"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"access_token", "access-control-allow-origin", "content-type"}

	r.Use(cors.New(config))
	mailer.InitDialer()
	utilities.InitRedis()
	models.DBConnect()
	route.SetupRoute(r)

	r.Run()
}

/*
mailer
multer
*/
