package main

import (
	"day03/config"
	route "day03/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	DB := config.InitDb()

	r := gin.Default()
	route.StudentRoutes(DB, r)
	r.Run(":80")
}
