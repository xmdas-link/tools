package encrypt

import (
	"crypto/sha256"
	"fmt"
)

type algSha256 struct {
}

func (s *algSha256) encode(content string) string {
	sum := sha256.Sum256([]byte(content))
	return fmt.Sprintf("%x", sum)
}

func (s *algSha256) hash(content string) []byte {
	hash := sha256.Sum256([]byte(content))
	return hash[:]
}
