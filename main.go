package main

import (
	"github.com/aldinofrizal/gin-ozamot-api/entity/models"
	"github.com/aldinofrizal/gin-ozamot-api/route"
	"github.com/aldinofrizal/gin-ozamot-api/services/mailer"
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
	models.DBConnect()
	route.SetupRoute(r)

	r.Run()
}

/*
mailer
multer
*/
