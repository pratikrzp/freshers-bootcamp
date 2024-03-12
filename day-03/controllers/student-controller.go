package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Working fine"})
}

func ReadStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Working fine"})
}

func ReadAllStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Working fine"})
}

func UpdateStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Working fine"})
}

func DeleteStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Working fine"})
}
