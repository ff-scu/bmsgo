package model

import "github.com/jinzhu/gorm"

// Book 定义用户model
type Book struct {
	gorm.Model
	BookName    string `json:"bookName"`
	BookAuthor  string `json:"bookAuth"`
	BookPublish string `json:"bookPublish"`
}
