package string_tool

import (
	"regexp"
)

func OnlyNumberWithLowercase(v string) bool {
	reg := regexp.MustCompile(`^[0-9a-z]+$`)
	return reg.MatchString(v)
}

// IsPhoneNumber check is a legal mobile phone number
func IsPhoneNumber(phone string) bool {
	reg := regexp.MustCompile("^[1][3456789][0-9]{9}$")
	return reg.MatchString(phone)
}
