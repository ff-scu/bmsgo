package controller

import (
	"bmsgo/common"
	"bmsgo/model"
	"bmsgo/response"
	"github.com/gin-gonic/gin"
)

func TestQtPost(ctx *gin.Context) {
	db := common.GetDB()
	var qtUser = model.QtUser{}
	_ = ctx.Bind(&qtUser)
	username := qtUser.UserName
	name := qtUser.Name
	sex := qtUser.Sex
	age := qtUser.Age

	newQtUser := model.QtUser{
		UserName: username,
		Name:     name,
		Sex:      sex,
		Age:      age,
	}
	db.Create(&newQtUser)
	//返回结果
	response.Success(ctx, gin.H{"data": "upload successful"}, "success")
}
//0x14 0x00 0x17 0x2d 0x01 0x6c 0x1e 0x1770
func GetNaviData(ctx *gin.Context)  {
	naviData:=gin.H{
		"data1":"1400172d016c1e1770",
	}
	response.SendData(ctx,naviData,"send success")
}
