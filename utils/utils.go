package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) []byte {
	hash := sha256.New()
	hash.Write(salt)
	hash.Write([]byte(password))
	return hash.Sum(nil)
}

func verifyPassword(storedHash, salt []byte, password string) bool {
	computedHash := hashPassword(password, salt)
	return bytes.Equal(storedHash, computedHash)
}
