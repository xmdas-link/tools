package string_tool

import (
	"errors"
	"strings"
)

/**
 * 格式化mac地址为标准:分割的mac
 */
func FormatMac(mac string) (string, error) {

	var builder strings.Builder

	mac = strings.ReplaceAll(mac, ":", "")
	mac = strings.ReplaceAll(mac, "-", "")
	mac = strings.ToUpper(mac)

	if len(mac) != 12 {
		return "", errors.New("MAC地址长度应为12个字符")
	}

	for i := 0; i < 10; i += 2 {
		builder.WriteString(mac[i : i+2])
		builder.WriteString(":")
	}
	builder.WriteString(mac[10:12])
	return builder.String(), nil
}
