package model

import "github.com/jinzhu/gorm"

type Paging struct {
	gorm.Model
	PageSize string `json:"page_size"`
	PageNum string `json:"page_num"`
	BookName string `json:"book_name"`
}