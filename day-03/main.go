package main

import (
	route "day03/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	route.StudentRoutes(r)
	r.Run(":80")
}
