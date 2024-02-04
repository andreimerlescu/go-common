package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(in string) (checksum string) {
	hash := sha256.New()
	hash.Write([]byte(in))
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}
