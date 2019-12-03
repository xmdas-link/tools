package string_tool

import (
	"math/rand"
	"time"
)

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
		//result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomNumber(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
		//result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomChar(length int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
		//result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
