package utils

import (
	"math/rand"
	"net/url"
	"regexp"
	"time"
)

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func EncodeBase62(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	enc := ""
	base := int64(len(base62Chars))

	for num > 0 {
		rem := num % base
		enc = string(base62Chars[rem]) + enc
		num /= base
	}

	return enc
}

func GenerateRandomShortKey(length int) string {
	shortKey := make([]byte, length)

	for idx := range shortKey {
		shortKey[idx] = base62Chars[rand.Intn(len(base62Chars))]
	}
	return string(shortKey)
}

func IsValidURL(inputURL string) bool {
	_, err := url.ParseRequestURI(inputURL)
	return err == nil
}

func IsValidShortKey(shortKey string) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9]{6,}$")
	return regex.MatchString(shortKey)
}
