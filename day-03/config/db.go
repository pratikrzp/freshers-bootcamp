package config

import (
	"day03/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dbUrl := "postgres://postgres:postgres@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}

	db.AutoMigrate(&models.Student{})

	return db
}
