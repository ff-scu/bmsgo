package controller

import (
	"bmsgo/common"
	"bmsgo/model"
	"bmsgo/response"
	"bmsgo/utils"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
	"time"
)

// Upload 上传文件到服务器
func Upload(ctx *gin.Context) {

	db := common.GetDB()

	form, err := ctx.MultipartForm()
	if err != nil {
		return
	}

	if len(form.File["file"]) <= 0 {
		response.Fail(ctx, gin.H{"file": "upload failed"}, "文件上传失败！")
		return
	}
	if form, err := ctx.MultipartForm(); err == nil {

		var fileInfo = &model.File{}

		//返回该次上传的文件名和文件地址的切片
		var fileAddressArr = make([]string, 0)
		var fileNameArr = make([]string, 0)

		//1.获取文件,返回的是一个切片
		files := form.File["file"]

		//2.循环全部的文件
		for _, file := range files {

			// 3.根据时间戳生成文件名
			fileNameInt := time.Now().Unix()
			fileNameStr := strconv.FormatInt(fileNameInt, 10)

			//4.新的文件名(如果是同时上传多张图片的时候就会同名，因此这里使用时间戳加文件名方式)
			fileName := fileNameStr + file.Filename

			//5.保存上传文件
			filePath := filepath.Join(utils.Mkdir("upload"), "/", fileName)
			err := ctx.SaveUploadedFile(file, filePath)
			if err != nil {
				return
			}
			filePath = filePath[7:]

			//将文件保存在数据库中
			fileInfo := model.File{
				FileName:    file.Filename,
				FileNewName: fileName,
				Address:     "localhost:8080/static/" + filePath,
				Size:        int(file.Size),
			}
			db.Create(&fileInfo)

			//将数据放入切片中
			fileAddressArr = append(fileAddressArr, fileInfo.Address)
			fileNameArr = append(fileNameArr, fileInfo.FileName)
		}

		response.Success(ctx, gin.H{"address": fileInfo.Address, "fileName": fileInfo.FileName}, "文件上传成功！")
		return
	} else {
		response.Fail(ctx, gin.H{"file": "upload failed"}, "文件上传失败！")
		return
	}
}

// SubmitForm 提交表单form到数据库中
func SubmitForm(ctx *gin.Context) {

	db := common.GetDB()

	var requestEmpInfo = model.EmployeeInfo{}

	_ = ctx.Bind(&requestEmpInfo)

	//获取参数
	newEmpInfo := &model.EmployeeInfo{
		Name:        requestEmpInfo.Name,
		Sex:         requestEmpInfo.Sex,
		Political:   requestEmpInfo.Political,
		Nation:      requestEmpInfo.Nation,
		Email:       int(requestEmpInfo.Email),
		Phone:       int(requestEmpInfo.Phone),
		Education:   requestEmpInfo.Education,
		School:      requestEmpInfo.School,
		Idcard:      int(requestEmpInfo.Idcard),
		Address:     requestEmpInfo.Address,
		Resource:    requestEmpInfo.Resource,
		Date1:       requestEmpInfo.Date1,
		Date2:       requestEmpInfo.Date2,
		FileaName:   requestEmpInfo.FileaName,
		FileAddress: requestEmpInfo.FileAddress,
	}

	db.Create(&newEmpInfo)

	response.Success(ctx, gin.H{"submitForm": "success"}, "表单提交成功")

}
