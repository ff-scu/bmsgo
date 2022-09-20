package dto

import "bmsgo/model"

type FileDto struct {
	FileName string
	Adrress  string
}

func ToFileDto(file model.File) FileDto {
	return FileDto{
		FileName: file.FileName,
		Adrress:  file.Address,
	}
}
