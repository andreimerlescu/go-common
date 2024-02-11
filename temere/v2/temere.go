package v2

import (
	"bytes"
	cr "crypto/rand"
	"math/big"
	mr "math/rand/v2"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func String(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, mr.Int32N(int32(n)), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = mr.Int32N(int32(n)), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func Bytes(n int) []byte {
	str := String(Integer(n))
	return bytes.NewBufferString(str).Bytes()
}

func Integer(limit int) int {
	c := make(chan int, 1)
	go func(l int) int {
		if l <= 0 {
			return 0
		}
		newInt, err := cr.Int(cr.Reader, big.NewInt(int64(l)))
		if err != nil {
			return l
		}
		data := int(newInt.Int64())
		c <- data
		return data
	}(limit)
	return <-c
}

func Integer64(limit int64) int64 {
	c := make(chan int64, 1)
	go func(l int64) int64 {
		if l <= 0 {
			return 0
		}
		newInt, err := cr.Int(cr.Reader, big.NewInt(int64(l)))
		if err != nil {
			return l
		}
		c <- newInt.Int64()
		return newInt.Int64()
	}(limit)
	return <-c
}

// RangeInteger return a random int that is between start and limit, will recursively run until match found
func RangeInteger(start int, limit int) int {
	i := Integer(limit)
	if i >= start && i <= limit {
		return i
	} else {
		return RangeInteger(start, limit) // retry
	}
}
