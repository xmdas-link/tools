package rsa

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"testing"
)

const privatePath = "D:\\IdeaProjects\\tomcatdemo\\src\\main\\resources\\private_key_pkcs8.pem"
const publicPath = "D:\\IdeaProjects\\tomcatdemo\\src\\main\\resources\\public_key_2048.pub"

func TestNewClient(t *testing.T) {

	client, err := NewClient(privatePath, publicPath)
	if err != nil {
		t.Error(err)
	}

	if client == nil {
		t.Error("新建加密client失败！")
	}

	// 加密文件
	text := "Hello world！";

	encryptText, err := client.Encrypt(text)
	t.Log("加密内容：", string(text))
	t.Log("秘文：", string(encryptText))

	decryptText, err := client.DecryptByte(encryptText)
	t.Log("解密后：", string(decryptText))

	if text != string(decryptText) {
		t.Error("加密失败")
	}

}

func TestJson(t *testing.T) {
	client, err := NewClient(privatePath, publicPath)
	if err != nil {
		t.Error(err)
	}

	if client == nil {
		t.Error("新建加密client失败！")
	}

	data := struct {
		CarNo string
		UserID uint
		GasName string
		GasCode string
		CardType string
	}{
		CarNo: "闽D233BH",
		UserID: 100234,
		GasName: "中石化",
		GasCode: "ZSH102301",
		CardType: "个人卡",
	}

	text, _ := json.Marshal(data)
	t.Log("需要加密内容：", string(text))

	encryptText, err := client.EncryptByte(text)
	t.Log("秘文：", string(encryptText))

	// 秘文写入文件（测试别的东西）
	f, _ := os.OpenFile("card_info_base.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)

	encodeString := base64.StdEncoding.EncodeToString(encryptText)

	f.Write([]byte(encodeString));

}
