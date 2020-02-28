package compress

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

type Compress interface {
	Zip(string) (string, error)   // 压缩 str -> zip -> base64_encode
	Unzip(string) (string, error) // 解压 str -> base64_decode -> unzip
}

type Gzip struct {
}

func (g Gzip) Zip(original string) (string, error) {
	if original == "" {
		return "", nil
	}

	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	w.Close()
	if _, err := w.Write([]byte(original)); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func (g Gzip) Unzip(original string) (string, error) {
	if original == "" {
		return "", nil
	}

	base64Dec, err := base64.StdEncoding.DecodeString(original)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(base64Dec)
	r, err := gzip.NewReader(buf)
	if err != nil {
		return "", err
	}
	defer r.Close()

	unzipContent, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(unzipContent), nil
}
