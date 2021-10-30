package utils

import (
	"crypto/md5"
	"strings"
)

// Md5Str 生成sessionID或密码
func Md5Str(str string) string {
	bt := md5.Sum([]byte(str))
	result := ""
	for _, i := range bt {
		result += string(i)
	}
	return strings.ToLower(result)
}
