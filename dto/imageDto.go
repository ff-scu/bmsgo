package dto

import (
	"bmsgo/model"
	"fmt"
)

// ImageDto 一组表格数据url
type ImageDto struct {
	Img0 string
	Img1 string
	Img2 string
	Img3 string
	Img4 string
}

// ToImageDto 返回image结构体中图片路径
func ToImageDto(imageArr []model.Image) []string {

	//var imageUrl []string

	sliceImageArr := make([]string, 0)

	for _, image := range imageArr {
		sliceImageArr = append(sliceImageArr, image.ImageUrl)
	}
	//fmt.Println("URL数组")
	//fmt.Println(sliceImageArr)

	return sliceImageArr
}

// ToImageDtoStruct 返回类型是数组包裹着结构体
func ToImageDtoStruct(imageArrString []string) []ImageDto {

	//要返回的数组，里面只有url
	sliceImageArrStruct := make([]ImageDto, 0)

	//fmt.Println("img0" + imageArrString[0])

	sliceImageStruct := ImageDto{
		Img0: imageArrString[0],
		Img1: imageArrString[1],
		Img2: imageArrString[0],
		Img3: imageArrString[0],
		Img4: imageArrString[0],
	}

	sliceImageArrStruct = append(sliceImageArrStruct, sliceImageStruct)

	fmt.Println(sliceImageArrStruct)

	return sliceImageArrStruct
}
