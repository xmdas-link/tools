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

// 文件扩展名检查
func IsFileNameType(fileName string, extends ...string) bool {

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

// 检查文件扩展名是否在允许范围，并返回
func CheckFileType(fileName string, extends ...string) string {
	var (
		fileArr = strings.Split(fileName, ".")
		extend  = strings.ToLower(fileArr[len(fileArr)-1])
	)

	for _, e := range extends {
		if extend == e {
			return extend
		}
	}

	return ""
}

// 是否word文件
func IsWordFile(file *multipart.FileHeader) bool {
	return IsFileNameType(file.Filename, TypeDoc, TypeDocx)
}

// 是否是Excel文件
func IsExcelFile(file *multipart.FileHeader) bool {
	return IsFileNameType(file.Filename, TypeXls, TypeXlsx)
}

// 是否是图片文件
func IsImageFile(file *multipart.FileHeader) bool {
	return IsFileNameType(file.Filename, TypePng, TypeJpg)
}

func IsPngFile(file *multipart.FileHeader) bool {
	return IsFileNameType(file.Filename, TypePng)
}

func IsJpgFile(file *multipart.FileHeader) bool {
	return IsFileNameType(file.Filename, TypeJpg)
}
