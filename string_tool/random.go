package string_tool

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	letterBytes       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits     = 6
	letterIdxMask     = 1<<letterIdxBits - 1
	letterIdxMax      = 63 / letterIdxBits
	numberBytes       = "0123456789"
	numberIdxBits     = 6
	numberIdxMask     = 1<<numberIdxBits - 1
	numberIdxMax      = (len(numberBytes) + 1) / numberIdxBits
	charBytes         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charIdxBits       = 6
	charIdxMask       = 1<<charIdxBits - 1
	charIdxMax        = (len(charBytes) + 1) / charIdxBits
	ISO_TIME_FORMAT_3 = "20060102150405.000000000"
)

func GetRandomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func GetRandomNumber(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), numberIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), numberIdxMax
		}
		if idx := int(cache & numberIdxMask); idx < len(numberBytes) {
			b[i] = numberBytes[idx]
			i--
		}
		cache >>= numberIdxBits
		remain--
	}
	return string(b)
}

func GetRandomChar(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), charIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), charIdxMax
		}
		if idx := int(cache & charIdxMask); idx < len(charBytes) {
			b[i] = charBytes[idx]
			i--
		}
		cache >>= charIdxBits
		remain--
	}
	return string(b)
}

func GetRandomWithTime(n int) string {
	return fmt.Sprint(time.Now().Format(ISO_TIME_FORMAT_3), GetRandomString(n))
}
