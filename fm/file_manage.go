package fm

import (
	"errors"
	"io"
	"mime/multipart"
)

var (
	fileManage FileManage
	basePath   string
)

// 文件管理
type FileManage interface {
	// 添加文件到文件路径
	AddFile(folder string, fileName string, file io.Reader) error
	// 读取web访问路径
	GetWebPath() string
	// 读取文件
	GetFile(folder string, fileName string) ([]byte, error)
}

func SetFileManage(m FileManage) {
	fileManage = m
}

func SetBasePath(b string) {
	basePath = b
}

func AddFileFromRequest(folder string, fileName string, file *multipart.FileHeader) error {
	if f, err := file.Open(); err != nil {
		return err
	} else {
		return AddFile(basePath+folder, fileName, f)
	}
}

func AddFile(folder string, fileName string, file io.Reader) error {
	if fileManage == nil {
		return errors.New("FileManage未初始化")
	}
	return fileManage.AddFile(basePath+folder, fileName, file)
}

func GetFile(folder string, fileName string) ([]byte, error) {
	if fileManage == nil {
		return nil, errors.New("FileManage未初始化")
	}
	return fileManage.GetFile(basePath+folder, fileName)

}

func GetFileString(folder string, fileName string) (string, error) {
	ret, err := GetFile(basePath+folder, fileName)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func GetWebPath() string {
	if fileManage != nil {
		return fileManage.GetWebPath() + basePath
	}
	return ""
}
