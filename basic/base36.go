package basic

import (
	"fmt"
	"strings"
)

/*

BASE 36
-------

WHY? It's case insensitive A-Z 0-9. That's why. Universally accepted because all systems can use the original intended string.

*/

const (
	IdentifierCharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func EncodeBase36(num int) string {
	if num == 0 {
		return "0"
	}

	var result strings.Builder
	for num > 0 {
		result.WriteString(string(IdentifierCharset[num%36]))
		num /= 36
	}

	return reverseString(result.String())
}

func DecodeBase36(s string) (int, error) {
	var num int
	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		if idx := strings.IndexRune(IdentifierCharset, char); idx >= 0 {
			num = num*36 + idx
		} else {
			// Handle invalid characters
			return 0, fmt.Errorf("invalid character: %c", char)
		}
	}
	return num, nil
}

func Encode64Base36(num int64) string {
	if num == 0 {
		return "0"
	}

	var result strings.Builder
	for num > 0 {
		result.WriteString(string(IdentifierCharset[num%36]))
		num /= 36
	}

	return reverseString(result.String())
}

func Decode64Base36(s string) (int64, error) {
	var num int64
	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		if idx := strings.IndexRune(IdentifierCharset, char); idx >= 0 {
			num = num*36 + int64(idx)
		} else {
			return 0, fmt.Errorf("invalid character: %c", char)
		}
	}
	return num, nil
}
