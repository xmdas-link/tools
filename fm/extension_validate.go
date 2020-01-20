package fm

import (
	"mime/multipart"
	"strings"
)

const (
	TypeDoc  = "doc"
	TypeDocx = "docx"
	TypePng  = "png"
	TypeJpg  = "jpeg"
	TypeXls  = "xls"
	TypeXlsx = "xlsx"
)

var ExtendsWord = []string{
	TypeDoc, TypeDocx,
}
var ExtendsImage = []string{
	TypePng, TypeJpg,
}
var ExtendsExcel = []string{
	TypeXls, TypeXlsx,
}

// 文件扩展名检查
func IsFileNameType(fileName string, extends []string) bool {

	var (
		fileArr = strings.Split(fileName, ".")
		extend  = strings.ToLower(fileArr[len(fileArr)-1])
	)

	for _, e := range extends {
		if extend == e {
			return true
		}
	}

	return false
}

// 是否word文件
func IsWordFile(file *multipart.FileHeader) bool {
	return IsFileNameType(file.Filename, ExtendsWord)
}

// 是否是Excel文件
func IsExcelFile(file *multipart.FileHeader) bool {
	return IsFileNameType(file.Filename, ExtendsExcel)
}

// 是否是图片文件
func IsImageFile(file *multipart.FileHeader) bool {
	return IsFileNameType(file.Filename, ExtendsImage)
}
