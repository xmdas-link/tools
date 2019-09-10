package string_tool

import "regexp"

func OnlyNumberWithLowercase(v string) bool {
	reg := regexp.MustCompile(`^[0-9a-z]+$`)
	return reg.MatchString(v)
}
