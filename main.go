package main

import (
	"bmsgo/common"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//加载配置项
	//InitConfig()
	//连接数据库
	db := common.InitDB()
	defer db.Close()

	//创建路由
	r := gin.Default()

	//绑定路由规则，执行函数
	r = CollectRoute(r)
	panic(r.Run())



}

func InitConfig() {
	//获取当前的工作目录
	workDir, _ := os.Getwd()

	config := viper.New()

	//设置读取的文件名
	config.SetConfigName("application")
	//设置读取的问价类型
	config.SetConfigType("yaml")
	//设置读取的文件路径
	config.AddConfigPath(workDir + "/config")
	fmt.Printf("get host%s \n", viper.GetString("datasource.host"))
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
