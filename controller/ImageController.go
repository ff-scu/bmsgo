package controller

import (
	"bmsgo/common"
	"bmsgo/dto"
	"bmsgo/model"
	"bmsgo/response"
	"bmsgo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// UploadImage 提交小臭宝图片到服务器 返回所有的图片url显示在表格中
func UploadImage(ctx *gin.Context) {
	db := common.GetDB()

	form, err := ctx.MultipartForm()
	if err != nil {
		return
	}

	if len(form.File["image"]) <= 0 {
		response.Fail(ctx, gin.H{"image": "upload failed"}, "图片上传失败！")
		return
	}

	if form, err := ctx.MultipartForm(); err == nil {

		//获取图片切片
		images := form.File["image"]

		//循环所有的图片
		for _, image := range images {

			//根据时间戳生成文件名
			imageInt := time.Now().Unix()
			imageStr := strconv.FormatInt(imageInt, 10)

			//新的图片的名字
			imageName := "图片" + imageStr + image.Filename

			//保存上传的图片
			imagePath := filepath.Join(utils.Mkdir("upload"), "/", imageName)
			err := ctx.SaveUploadedFile(image, imagePath)
			if err != nil {
				return
			}
			imagePath = imagePath[7:]

			//将图片保存在数据库中
			newImage := model.Image{
				Name:     image.Filename,
				ImageUrl: "http://localhost:8080/static/" + imagePath,
				Describe: "",
			}

			db.Create(&newImage)

			response.Success(ctx, gin.H{"imageUrl": newImage.ImageUrl}, "图片上传成功！")
		}
	}
}

// ShowImage 请求所有的图片数据显示出来 返回为图片的地址数组 类似一个分页操作，取两个数组，每个数组5个元素
func ShowImage(ctx *gin.Context) {

	db := common.GetDB()

	//imageArr 返回的所有数据
	var imageArr []model.Image

	db.Find(&imageArr)

	//将url放到数组中
	imageUrlArr := dto.ToImageDto(imageArr)

	//将数组中的url放到结构体中，封装为结构体数组
	sliceImageArrStruct := make([]dto.ImageDto, 0)

	var sliceImageStruct dto.ImageDto

	if len(imageUrlArr) <= 5 { //考虑数组越界问题
		if len(imageUrlArr) == 1 {
			sliceImageStruct = dto.ImageDto{
				Img0: imageUrlArr[0],
			}
		}
		if len(imageUrlArr) == 2 {
			sliceImageStruct = dto.ImageDto{
				Img0: imageUrlArr[0],
				Img1: imageUrlArr[1],
			}
		}
		if len(imageUrlArr) == 3 {
			sliceImageStruct = dto.ImageDto{
				Img0: imageUrlArr[0],
				Img1: imageUrlArr[1],
				Img2: imageUrlArr[2],
			}
		}
		if len(imageUrlArr) == 4 {
			sliceImageStruct = dto.ImageDto{
				Img0: imageUrlArr[0],
				Img1: imageUrlArr[1],
				Img2: imageUrlArr[2],
				Img3: imageUrlArr[3],
			}
		}
		if len(imageUrlArr) == 5 {
			sliceImageStruct = dto.ImageDto{
				Img0: imageUrlArr[0],
				Img1: imageUrlArr[1],
				Img2: imageUrlArr[2],
				Img3: imageUrlArr[3],
				Img4: imageUrlArr[4],
			}
		}
		sliceImageArrStruct = append(sliceImageArrStruct, sliceImageStruct)
	} else {
		//len>5
		var cycleNum int //循环次数

		var x = len(imageUrlArr) //数组长度

		cycleNum = x / 5

		for i := 0; i < cycleNum; i++ {
			sliceImageStruct = dto.ImageDto{
				Img0: imageUrlArr[0+i*5],
				Img1: imageUrlArr[1+i*5],
				Img2: imageUrlArr[2+i*5],
				Img3: imageUrlArr[3+i*5],
				Img4: imageUrlArr[4+i*5],
			}
			sliceImageArrStruct = append(sliceImageArrStruct, sliceImageStruct)
		}

		if x%5 != 0 {
			if len(imageUrlArr)%5 == 1 {
				sliceImageStruct = dto.ImageDto{
					Img0: imageUrlArr[0+cycleNum*5],
				}
			}
			if len(imageUrlArr)%5 == 2 {
				sliceImageStruct = dto.ImageDto{
					Img0: imageUrlArr[0+cycleNum*5],
					Img1: imageUrlArr[1+cycleNum*5],
				}
			}
			if len(imageUrlArr)%5 == 3 {
				sliceImageStruct = dto.ImageDto{
					Img0: imageUrlArr[0],
					Img1: imageUrlArr[1+cycleNum*5],
					Img2: imageUrlArr[2+cycleNum*5],
				}
			}
			if len(imageUrlArr)%5 == 4 {
				sliceImageStruct = dto.ImageDto{
					Img0: imageUrlArr[0+cycleNum*5],
					Img1: imageUrlArr[1+cycleNum*5],
					Img2: imageUrlArr[2+cycleNum*5],
					Img3: imageUrlArr[3+cycleNum*5],
				}
			}
			sliceImageArrStruct = append(sliceImageArrStruct, sliceImageStruct)
		}
	}

	response.Success(ctx, gin.H{"imageUrl": sliceImageArrStruct, "imageUrlArr": imageUrlArr}, "返回的图片地址数组")

}

//删除照片
func DeleteImage(ctx *gin.Context) {
	db := common.GetDB()

	var requestImage model.Image

	_ = ctx.Bind(&requestImage)

	imagUrl := requestImage.ImageUrl

	db.Where("image_url = ?", imagUrl).Delete(&requestImage)

	//path:=requestImage.ImageUrl  "upload/2022/07/03/图片1656778395臭宝.jpg"

	path:="upload/"+requestImage.ImageUrl[29:]

	err:= os.Remove(path)
	if err !=nil {
		fmt.Println(err)
	}else {
		fmt.Println("删除成功")
	}

	response.Success(ctx, nil, "删除图片成功")
}
