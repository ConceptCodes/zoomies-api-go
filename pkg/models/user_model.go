package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primary_key" json:"id"`
	Email     string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string `gorm:"size:255" json:"password"`
	FirstName string `gorm:"size:255" json:"firstName"`
	LastName  string `gorm:"size:255" json:"lastName"`
}

type SimpleUser struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
