package hash

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
)

type PasswordHasher interface {
	GenerateHashedPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}

type Hash struct {
	salt string
}

func New(salt string) *Hash {
	return &Hash{salt}
}

func (h *Hash) GenerateHashedPassword(password string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}

func (h *Hash) CompareHashAndPassword(hashedPassword, password string) error {
	hashedClientPassword, err := h.GenerateHashedPassword(password)
	if err != nil {
		return err
	}

	res := strings.Compare(hashedPassword, hashedClientPassword)
	if res != 0 {
		return errors.New("passwords do not match")
	}

	return nil
}
