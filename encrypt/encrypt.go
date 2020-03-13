package encrypt

import (
	"fmt"
	"strings"
)

const (
	EncryptAlgSha256 string = "SHA256"
	EncryptAlgSm3    string = "SM3"
)

var SupportEncryptAlg = map[string]bool{
	EncryptAlgSha256: true,
	EncryptAlgSm3:    true,
}

type Encrypt struct {
	Alg     string // 加密算法
	Content string
}

func NewEncrypt(alg, content string) (*Encrypt, error) {
	alg = strings.ToUpper(alg)
	if _, ok := SupportEncryptAlg[alg]; !ok {
		return nil, fmt.Errorf("not supoort encrypt alg: %v", alg)
	}

	return &Encrypt{
		Alg:     alg,
		Content: content,
	}, nil
}

func (e *Encrypt) Crypto() (string, error) {
	if len(e.Content) == 0 {
		return "", nil
	}

	alg, err := e.algFactory()
	if err != nil {
		return "", err
	}

	return alg.encode(e.Content), nil
}

func (e *Encrypt) algFactory() (encrypt, error) {
	switch e.Alg {
	case EncryptAlgSha256:
		return &algSha256{}, nil
	case EncryptAlgSm3:
		return &algSm3{}, nil
	default:
		return nil, fmt.Errorf("unknowed encrypt alg: %v", e.Alg)
	}
}
