package model

import "github.com/jinzhu/gorm"

// User 定义用户model
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Password  string `gorm:"size:255;not null"`
	Telephone string
}
