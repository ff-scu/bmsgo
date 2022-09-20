package model

import "github.com/jinzhu/gorm"

// File 定义文件
type File struct {
	gorm.Model

	FileName    string //文件名
	FileNewName string
	Address     string //文件地址
	Size        int    //文件大小
}
