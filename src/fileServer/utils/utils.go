package utils

import (
	"crypto/sha1"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

func GenerateLotToken(token string) string {
	return token // потом можно добавить логику создания замороченного токена
}

func GenSha1(input string) string {
	h := sha1.New()
	if _, err := io.WriteString(h, input); err != nil {
		logrus.Errorf("Error hashing %v string: %v", input, err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
