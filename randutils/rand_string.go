package randutils

import (
	"math/rand"
	"time"
	"unsafe"
)

const PoolAlphaUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const PoolAlphaLowercase = "abcdefghijklmnopqrstuvwxyz"
const PoolAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const PoolAlphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandString Produces a random string with length stringSize.
// 	The output string will only contain the basic uppercase and
//	lowercase english letters.
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandString(stringSize int) string {
	return RandStringWithPool(stringSize, PoolAlpha)
}

func RandStringWithPool(stringSize int, letterBytes string) string {
	b := make([]byte, stringSize)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := stringSize-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
