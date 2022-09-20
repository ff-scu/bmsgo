package model

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	Name     string
	ImageUrl string
	Describe string
}
