package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabetic = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt 返回min到max之间的一个随机数
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString 生成一个长度为n的随机字符串
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabetic)
	for i := 0; i < n; i++ {
		c := alphabetic[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencys := []string{"USD", "EUR", "GBP", "RMB"}
	n := len(currencys)
	return currencys[rand.Intn(n)]
}
