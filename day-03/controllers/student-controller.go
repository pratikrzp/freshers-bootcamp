package controllers

import (
	"day03/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

func NewStudentController(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (h StudentHandler) CreateStudent(c *gin.Context) {
	student := models.Student{}
	if err := c.BindJSON(&student); err != nil {
		panic("Failed to bind payload")
	}
	result := h.DB.Create(&student)
	if result.Error != nil {
		panic("Failed to create student")
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Student created"})
}

func (h StudentHandler) ReadStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Working fine"})
}

func (h StudentHandler) ReadAllStudent(c *gin.Context) {
	student := models.Student{}
	h.DB.Find(&student)
	c.JSON(http.StatusOK, gin.H{"message": "Working fine", "data": student})
}

func (h StudentHandler) UpdateStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Working fine"})
}

func (h StudentHandler) DeleteStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Working fine"})
}
