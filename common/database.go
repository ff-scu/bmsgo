package common

import (
	"bmsgo/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/url"
)

//全局的db
var DB *gorm.DB

//开启数据库连接
func InitDB() *gorm.DB {
	//driverName:=viper.GetString("datasource.driverName")
	//host:=viper.GetString("datasource.host")
	//port:=viper.GetString("datasource.port")
	//database:=viper.GetString("datasource.database")
	//username:=viper.GetString("datasource.username")
	//password:=viper.GetString("datasource.password")
	//charset:=viper.GetString("datasource.charset")
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "test"
	username := "root"
	password := "1234"
	charset := "utf8"
	loc := "Asia/Shanghai"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	//自动创建数据表
	db.AutoMigrate(&model.User{})

	db.AutoMigrate(&model.Book{})

	db.AutoMigrate(&model.File{})

	db.AutoMigrate(&model.EmployeeInfo{})

	db.AutoMigrate(&model.Image{})

	db.AutoMigrate(&model.QtUser{})

	//赋值
	DB = db
	return db
}

//写一个方法获取db实例
func GetDB() *gorm.DB {
	return DB
}
