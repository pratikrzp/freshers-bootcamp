package route

import (
	"day03/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine) {
	r1 := router.Group("/student", func(ctx *gin.Context) {
		fmt.Println("New request received with student group")
	})
	{
		r1.GET("/", controllers.ReadAllStudent)
		r1.GET("/:id", controllers.ReadStudent)
		r1.POST("/", controllers.CreateStudent)
		r1.PUT("/:id", controllers.UpdateStudent)
		r1.DELETE("/:id", controllers.DeleteStudent)
	}
}
