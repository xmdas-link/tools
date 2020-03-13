package encrypt

import (
	"fmt"

	"github.com/ZZMarquis/gm/sm3"
)

type algSm3 struct {
}

func (s *algSm3) encode(content string) string {
	sum := sm3.Sum([]byte(content))
	return fmt.Sprintf("%x", sum)
}
