package rand

import (
	crand "crypto/rand"
	"math/rand"
	"strconv"
	"time"
)

const (
	ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ALPHANUM = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ALPHASYM = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz~!@#$%^&*_+?-="
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString 获取随机字符串
func RandomString(count int, sets ...byte) string {
	if count == 0 {
		return ""
	}
	if len(sets) == 0 {
		sets = []byte(ALPHANUM)
	}

	var ret = make([]byte, count)
	setlen := len(ALPHANUM)
	for i := 0; i < count; i++ {
		ret[i] = sets[RandomInt(0, setlen-1)]
	}
	return string(ret)
}

// RandomBytes 获取随机字符
func RandomBytes(count int) []byte {
	b := make([]byte, count)
	crand.Read(b)
	return b
}

func RandomInt(start, end int) int {
	return rand.Intn(end-start+1) + start
}

func RandomInt64(start, end int64) int64 {
	return rand.Int63n(end-start+1) + start
}

func RandomCode(sum int) (code string) {
	for i := 0; i < sum; i++ {
		code = code + strconv.Itoa(rand.Intn(9))
	}
	return code
}
