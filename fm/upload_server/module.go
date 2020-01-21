package upload_server

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

/**
 * 上传到116上的image server，一个非常简单的服务器
 */
type Module struct {
	Url       string
	VisitHost string
}

func (m *Module) GetWebPath() string {
	return m.VisitHost
}

func (m *Module) AddFile(folder string, fileName string, file io.Reader) error {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return err
	}

	writer.WriteField("folder", folder)
	writer.WriteField("file_name", fileName)

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", m.Url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ret := &UploadResult{}
	if err := json.Unmarshal(content, ret); err != nil {
		return err
	}

	if ret.Ret != 1 {
		return errors.New(ret.Error)
	}

	return nil
}

func (m *Module) GetFile(folder string, fileName string) ([]byte, error) {

	// 实际url地址
	var fileUrl = m.VisitHost + folder + "/" + fileName
	resp, err := http.Get(fileUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("获取解析文件失败")
	}

	return ioutil.ReadAll(resp.Body)

}

func New(url string, vHost string) *Module {
	m := &Module{
		Url:       url,
		VisitHost: vHost,
	}
	return m
}
