package string_tool

func AddMosaicForString(s string) string {
	r := []rune(s)
	var result string
	switch {
	case len(r) > 10:
		result = string(r[:3]) + "****" + string(r[len(r)-3:])
	case len(r) > 5:
		result = string(r[:2]) + "***" + string(r[len(r)-2:])
	case len(r) > 2:
		result = string(r[:1]) + "**" + string(r[len(r)-1:])
	case len(r) == 2:
		result = string(r[:1]) + "*"
	case len(r) == 1:
		result = "*"
	}
	return result
}
