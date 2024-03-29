package upload_server

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

/**
 * 上传到116上的image server，一个非常简单的服务器
 */
type Module struct {
	Url       string
	CopyUrl   string
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
		Timeout: 30 * time.Second,
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
		return nil, errors.New("获取解析文件失败：" + fileUrl)
	}

	return ioutil.ReadAll(resp.Body)

}

func (m *Module) CopyFile(fromPath string, toPath string, newFileName string) error {
	if len(m.CopyUrl) == 0 {
		return errors.New("未配置CopyUrl")
	}

	var (
		params = url.Values{}
	)

	params.Add("from", fromPath)
	params.Add("to", toPath)
	params.Add("file_name", newFileName)

	resp, err := http.PostForm(m.CopyUrl, params)
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

func (m *Module) UploadTempFile(fileName string, file io.Reader) (download string, err error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return "", err
	}

	writer.WriteField("file_name", fileName)

	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", m.Url, body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ret := &UploadResult{}
	if err := json.Unmarshal(content, ret); err != nil {
		return "", err
	}

	if ret.Ret != 1 {
		return "", errors.New(ret.Error)
	}

	return ret.Url, nil
}

func New(url string, vHost string) *Module {
	m := &Module{
		Url:       url,
		VisitHost: vHost,
	}
	return m
}
