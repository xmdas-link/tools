package string_tool

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
