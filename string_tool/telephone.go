package string_tool

import "strings"

func FormatChinaTelephone(chinaTelephone string) string {
	return strings.Replace(chinaTelephone, "86-", "", 1)
}

func GetChinaTelephone(telephone string) string {
	return "86-" + telephone
}
