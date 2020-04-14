package fm

import (
	"io/ioutil"
)

// FileCopy copy file from src to dst
func FileCopy(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, input, 0777)
	if err != nil {
		return err
	}

	return nil
}
