package fm

import (
	"errors"
	"io"
	"mime/multipart"
)

var (
	fileManage FileManage
)

// 文件管理
type FileManage interface {
	// 添加文件到文件路径
	AddFile(folder string, fileName string, file io.Reader) error
	// 读取web访问路径
	GetWebPath() string
}

func SetFileManage(m FileManage) {
	fileManage = m
}

func AddFileFromRequest(folder string, fileName string, file *multipart.FileHeader) error {
	if f, err := file.Open(); err != nil {
		return err
	} else {
		return AddFile(folder, fileName, f)
	}
}

func AddFile(folder string, fileName string, file io.Reader) error {
	if fileManage == nil {
		return errors.New("FileManage未初始化")
	}
	return fileManage.AddFile(folder, fileName, file)
}

func GetWebPath() string {
	if fileManage != nil {
		return fileManage.GetWebPath()
	}
	return ""
}
