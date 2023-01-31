package hash

import (
	"crypto/sha256"
	"fmt"
)

type PasswordHasher interface {
	GenerateEncryptedPassword(password string) (string, error)
}

type Hash struct {
	salt string
}

func New(salt string) *Hash {
	return &Hash{salt: salt}
}

func (h *Hash) GenerateEncryptedPassword(password string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}
