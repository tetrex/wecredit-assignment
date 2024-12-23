package otp

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewOtp(length int) (string, error) {
	result := make([]byte, length)
	charsetLen := len(charset)

	for i := range result {
		randomByte, err := rand.Int(rand.Reader, big.NewInt(int64(charsetLen)))
		if err != nil {
			return "", err
		}
		result[i] = charset[randomByte.Int64()]
	}

	return string(result), nil
}
