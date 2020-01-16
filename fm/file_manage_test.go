package fm

import (
	"github.com/gin-gonic/gin"
	"github.com/xmdas-link/tools/fm/upload_server"
	"os"
	"testing"
)

func TestAddFile(t *testing.T) {

	var (
		url = "http://192.168.0.116:11801/upload"
		web = "http://192.168.0.116:1180/"
	)

	// 设置文件上传配置
	uploadServer := upload_server.New(url, web)
	SetFileManage(uploadServer)

	// 从本地读取文件
	file, _ := os.Open("文件地址")

	// 添加文件
	err := AddFile("test/2019/", "text.log", file)
	if err != nil {
		t.Error(err)
	}

}

func TestAddFileFromRequest(t *testing.T) {
	var (
		url = "http://192.168.0.116:11801/upload"
		web = "http://192.168.0.116:1180/"
		c   *gin.Context
	)

	// 设置文件上传配置
	uploadServer := upload_server.New(url, web)
	SetFileManage(uploadServer)

	// 从gin的Request上获取multipart/form-data
	file, _ := c.FormFile("file")
	// 添加文件
	err := AddFileFromRequest("test/2019/", "text.log", file)
	if err != nil {
		t.Error(err)
	}

}
