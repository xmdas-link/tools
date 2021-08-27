package rsa

import "crypto"

type Cipher interface {
	// 加密
	Encrypt(text string) ([]byte, error)
	EncryptByte(text []byte) ([]byte, error)
	// 解密
	Decrypt(text string) ([]byte, error)
	DecryptByte(text []byte) ([]byte, error)
	// 签名
	Sign(src []byte, hash crypto.Hash) ([]byte, error)
	// 验签
	Verify(src []byte, sign []byte, hash crypto.Hash) error
}