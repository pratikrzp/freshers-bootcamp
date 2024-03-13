package route

import (
	"day03/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StudentRoutes(DB *gorm.DB, router *gin.Engine) {
	studentHandler := controllers.NewStudentController(DB)
	r1 := router.Group("/student", func(ctx *gin.Context) {
		fmt.Println("New request received with student group")
	})
	{
		r1.GET("/", studentHandler.ReadAllStudent)
		r1.GET("/:id", studentHandler.ReadStudent)
		r1.POST("/", studentHandler.CreateStudent)
		r1.PUT("/:id", studentHandler.UpdateStudent)
		r1.DELETE("/:id", studentHandler.DeleteStudent)
	}
}
