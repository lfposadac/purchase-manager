package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"uniqueIndex"`
	Role   string `json:"role"` //applicant or financier
	Active bool   `json:"active" gorm:"default:true"`
}
