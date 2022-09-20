package model

import "github.com/jinzhu/gorm"

type ImageTable struct {
	gorm.Model
	Img0 string
	Img1 string
	Img2 string
	Img3 string
	Img4 string
}
