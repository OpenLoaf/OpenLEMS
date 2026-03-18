package c_util

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// IsPasswordHash reports whether the stored value looks like a bcrypt hash.
func IsPasswordHash(value string) bool {
	return strings.HasPrefix(value, "$2a$") ||
		strings.HasPrefix(value, "$2b$") ||
		strings.HasPrefix(value, "$2y$")
}

// HashPassword converts a plaintext password into a bcrypt hash.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// VerifyPassword supports both bcrypt hashes and legacy plaintext values.
func VerifyPassword(storedValue string, plaintext string) (bool, error) {
	if storedValue == "" {
		return false, nil
	}

	if !IsPasswordHash(storedValue) {
		return storedValue == plaintext, nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedValue), []byte(plaintext))
	if err == nil {
		return true, nil
	}
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}
	return false, err
}
