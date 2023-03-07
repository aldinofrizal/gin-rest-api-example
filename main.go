package main

import (
	"github.com/aldinofrizal/gin-ozamot-api/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	route.SetupRoute(r)

	r.Run()
}
