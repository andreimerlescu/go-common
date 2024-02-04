package temere

import (
	"bytes"
	cr "crypto/rand"
	"math/big"
	mr "math/rand"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var randSource = mr.NewSource(time.Now().UnixNano())

func String(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, randSource.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSource.Int63(), letterIdxMax
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

func Bytes() []byte {
	str := String(Integer(999))
	return bytes.NewBufferString(str).Bytes()
}

func Integer(limit int) int {
	c := make(chan int, 1)
	go func(l int) int {
		// Seed the rand engine
		mr.New( // Who?
			mr.NewSource( // Who?
				time.Now().UnixNano())) // When? LOL

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

// RangeInteger return a random int that is between start and limit, will recursively run until match found
func RangeInteger(start int, limit int) int {
	i := Integer(limit)
	if i >= start && i <= limit {
		return i
	} else {
		return RangeInteger(start, limit) // retry
	}
}
