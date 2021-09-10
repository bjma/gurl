package utils

import (
	"strconv"
)

func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

func ByteArrayToStr(b []byte) string {
	return string(b)
}

func StrToByteArray(s string) []byte {
	return []byte(s)
}
