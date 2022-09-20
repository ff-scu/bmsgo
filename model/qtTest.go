package model

import "github.com/jinzhu/gorm"

type QtUser struct {
	gorm.Model
	UserName string
	Name string
	Sex string
	Age string
}
