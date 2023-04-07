package hash

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/v1adhope/crypto-diary/internal/entity"
	"golang.org/x/crypto/scrypt"
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
	dk, err := scrypt.Key([]byte(password), []byte(h.salt), 1<<15, 8, 1, 32)
	if err != nil {
		return "", fmt.Errorf("hash: GenerateHashedPassword: Write: %w", err)
	}

	return base64.StdEncoding.EncodeToString(dk), nil
}

func (h *Hash) CompareHashAndPassword(hashedPassword, password string) error {
	hashedClientPassword, err := h.GenerateHashedPassword(password)
	if err != nil {
		return fmt.Errorf("hash: CompareHashAndPassword: GenerateHashedPassword: %w", err)
	}

	match := strings.Compare(hashedPassword, hashedClientPassword)
	if match != 0 {
		return fmt.Errorf("hash: CompareHashAndPassword: Compare: %w", entity.ErrWrongPassword)
	}

	return nil
}
