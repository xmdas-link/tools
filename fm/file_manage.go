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
	// 复制文件
	CopyFile(fromPath string, toPath string, newFileName string) error
}

func SetFileManage(m FileManage) {
	fileManage = m
}

func SetBasePath(b string) {
	basePath = b
}

/**
 * 从gin的请求中直接拿file文件
 */
func AddFileFromRequest(folder string, fileName string, file *multipart.FileHeader) error {
	if f, err := file.Open(); err != nil {
		return err
	} else {
		return AddFile(folder, fileName, f)
	}
}

/**
 * 从io中拿file文件
 */
func AddFile(folder string, fileName string, file io.Reader) error {
	if fileManage == nil {
		return errors.New("FileManage未初始化")
	}
	return fileManage.AddFile(basePath+folder, fileName, file)
}

/**
 * 获取文件
 */
func GetFile(folder string, fileName string) ([]byte, error) {
	if fileManage == nil {
		return nil, errors.New("FileManage未初始化")
	}
	return fileManage.GetFile(basePath+folder, fileName)

}

/**
 * 以字符串的格式获取文件
 */
func GetFileString(folder string, fileName string) (string, error) {
	ret, err := GetFile(folder, fileName)
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

func CopyFile(fromPath string, toPath string, newFileName string) error {
	if fileManage == nil {
		return errors.New("FileManage未初始化")
	}
	return fileManage.CopyFile(basePath+fromPath, basePath+toPath, newFileName)
}
