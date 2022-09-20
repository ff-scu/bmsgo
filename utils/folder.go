package utils

import (
	"os"
	"path/filepath"
	"time"
)

// Mkdir 定义一个创建文件目录的方法
func Mkdir(basePath string) string {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(basePath, folderName)
	//使用mkdirall会创建多层级目录
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return ""
	}
	return folderPath
}
