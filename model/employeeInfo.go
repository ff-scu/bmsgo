package model

import "github.com/jinzhu/gorm"

// EmployeeInfo 定义上传的员工信息
type EmployeeInfo struct {
	gorm.Model
	Name        string
	Sex         string
	Political   string
	Nation      string
	Email       int
	Phone       int
	Education   string
	School      string
	Idcard      int
	Address     string
	Resource    string
	Date1       string
	Date2       string
	FileaName   string
	FileAddress string
}
