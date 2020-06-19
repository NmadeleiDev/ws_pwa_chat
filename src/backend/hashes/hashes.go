package hashes

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func CalculateSha256(value string) string {
	sha256Calculate := sha256.Sum256([]byte(value))
	hash := fmt.Sprintf("%x", sha256Calculate)
	return hash
}

func CalculateSha1(value string) string {
	sha256Calculate := sha1.Sum([]byte(value))
	hash := fmt.Sprintf("%x", sha256Calculate)
	return hash
}
