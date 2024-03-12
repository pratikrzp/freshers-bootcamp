package route

import (
	"day03/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine) {
	r1 := router.Group("/user")
	r1.GET("/", controllers.ReadAllStudent)
	r1.GET("/:id", controllers.ReadStudent)
	r1.POST("/", controllers.CreateStudent)
	r1.PUT("/:id", controllers.UpdateStudent)
	r1.DELETE("/:id", controllers.DeleteStudent)
}
