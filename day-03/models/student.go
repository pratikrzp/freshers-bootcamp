package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	DOB      string `json:"dob"`
}
